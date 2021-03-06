package result

import (
	"bytes"
	"github.com/android-test-runner/atr/devices"
	"github.com/android-test-runner/atr/files"
	"html/template"
	"sort"
)

const cssTemplate = `
* {
	font-family: Arial;
}
pre {
	font-family: monospace;
}
p.title {
	margin: 0px;
	font-weight: bold;
}
video {
	width: 400px;
	padding: 5px;
}
ul.testResults {
	list-style-type:none;
	padding-left: 0;
}
li.testResult {
	padding-left: 5px;
}
ul.extras {
	padding-bottom: 5px;
}
.overviewCount {
	font-weight: bold;
	display: inline;
}

.Passed {
	border-left: 5px solid green;
	border-bottom: 1px solid green;
}
.numPassed {
	color: green;
}
.Failed {
	border-left: 5px solid red;
	border-bottom: 1px solid red;
}
.numFailed {
	color: red;
}
.Errored {
	border-left: 5px solid red;
	border-bottom: 1px solid red;
}
.Skipped {
	border-left: 5px solid yellow;
	border-bottom: 1px solid yellow;
}
.numSkipped {
	color: orange;
}
`

const htmlTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<title>ATR Results</title>
		<link href="{{ .ResultsCss }}" rel="stylesheet" />
	</head>
	<body>
		<h1>Overview</h1>
		<ul>
		{{ range $testResult := .Results }}
			<li><a href="#{{ $testResult.DeviceName }}">{{ $testResult.DeviceName }}</a>: 
				{{ if $testResult.NumFailed }}
					<p class="overviewCount numFailed">{{ $testResult.NumFailed }} Failed</p>, 
				{{ end }}
				{{ if $testResult.NumSkipped }}
					<p class="overviewCount numSkipped">{{ $testResult.NumSkipped }} Skipped</p>,
				{{ end }}
				<p class="overviewCount numPassed">{{ $testResult.NumPassed }} Passed</p> 
			</li>
		{{ end }}
		</ul>

		{{ range $testResult := .Results }}
			<h1><a id="{{ $testResult.DeviceName }}"/>{{ $testResult.DeviceName }}</h1>
			<ul class="testResults">
			{{ range $result := $testResult.Results }}
				<li class="testResult {{ $result.Status }}">
				<p class="title">{{ $result.TestName }}: {{$result.Status}}</p>
				{{ if $result.Output }}
					<pre>{{ $result.Output }}</pre>
				{{ end }}
				{{ if $result.Video }}
					<a href="{{$result.Video}}">Video</a></br>
					<video controls>
						<source src="{{$result.Video}}" type="video/mp4" />
						Your browser does not support the video tag.
					</video>
				{{ end }}
				<ul class="extras">
					{{ range $extra := $result.Extras }}
						<li><a href="{{ $extra.Link }}">{{ $extra.Name }}</a></li>
					{{ end }}
				</ul>
				</li>
			{{ end }}
			</ul>
		{{ end }}
	</body>
</html>
`

type outputHtml struct {
	Results    []resultsForDeviceHtml
	ResultsCss string
}

type resultsForDeviceHtml struct {
	DeviceName string
	Results    []resultHtml
	NumFailed  int
	NumPassed  int
	NumSkipped int
}

type resultHtml struct {
	TestName  string
	Status    string
	IsFailure bool
	Output    string
	Video     string
	Extras    []extraHtml
}

type extraHtml struct {
	Name string
	Link string
}

type HtmlFormatter interface {
	FormatResults(map[devices.Device]TestResults) ([]files.File, error)
}

type htmlFormatterImpl struct{}

func NewHtmlFormatter() HtmlFormatter {
	return htmlFormatterImpl{}
}

func (formatter htmlFormatterImpl) FormatResults(resultsByDevice map[devices.Device]TestResults) ([]files.File, error) {
	parsedTemplate, templateError := template.New("html").Parse(htmlTemplate)
	if templateError != nil {
		return []files.File{}, templateError
	}

	cssFileName := "results.css"
	var content bytes.Buffer
	htmlOutput := formatter.toHtmlOutput(resultsByDevice, cssFileName)
	templateExecutionError := parsedTemplate.Execute(&content, htmlOutput)
	if templateExecutionError != nil {
		return []files.File{}, templateExecutionError
	}

	htmlFile := files.File{
		Name:    "results.html",
		Content: content.String(),
	}

	cssFile := files.File{
		Name:    cssFileName,
		Content: cssTemplate,
	}

	return []files.File{htmlFile, cssFile}, nil
}

func (formatter htmlFormatterImpl) toHtmlOutput(resultsByDevice map[devices.Device]TestResults, cssFileName string) outputHtml {
	resultsForDeviceHtmls := []resultsForDeviceHtml{}
	for device, testResults := range resultsByDevice {
		resultsHtml := []resultHtml{}
		numFailed := 0
		numPassed := 0
		numSkipped := 0
		for _, result := range testResults.Results {
			if result.IsFailure() {
				numFailed += 1
			} else if result.IsSkipped() {
				numSkipped += 1
			} else {
				numPassed += 1
			}
			resultsHtml = append(resultsHtml, toHtmlResult(result))
		}
		sortResultHtmlsByStatus(resultsHtml)

		resultsAndDevice := resultsForDeviceHtml{
			DeviceName: device.Serial,
			Results:    resultsHtml,
			NumFailed:  numFailed,
			NumPassed:  numPassed,
			NumSkipped: numSkipped,
		}

		resultsForDeviceHtmls = append(resultsForDeviceHtmls, resultsAndDevice)
	}

	return outputHtml{Results: resultsForDeviceHtmls, ResultsCss: cssFileName}
}

func toHtmlResult(result Result) resultHtml {
	htmlExtras := []extraHtml{}
	video := ""
	for _, extra := range result.Extras {
		if extra.Type == File {
			htmlExtras = append(htmlExtras, toHtmlExtra(extra))
		} else if extra.Type == Video {
			video = extra.Value
		}
	}

	output := ""
	if result.IsFailure() {
		output = result.Output
	}

	return resultHtml{
		TestName:  result.Test.FullName(),
		Status:    result.Status.toString(),
		IsFailure: result.IsFailure(),
		Output:    output,
		Video:     video,
		Extras:    htmlExtras,
	}
}

func toHtmlExtra(extra Extra) extraHtml {
	return extraHtml{
		Name: extra.Name,
		Link: extra.Value,
	}
}

type ByStatus []resultHtml

func (s ByStatus) Len() int           { return len(s) }
func (s ByStatus) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByStatus) Less(i, j int) bool { return toSortNumber(s[i]) < toSortNumber(s[j]) }

func sortResultHtmlsByStatus(resultHtmls []resultHtml) {
	sort.Sort(ByStatus(resultHtmls))
}

func toSortNumber(result resultHtml) int {
	// A failed test result shall be before a successful one
	if result.IsFailure {
		return 0
	} else {
		return 1
	}
}
