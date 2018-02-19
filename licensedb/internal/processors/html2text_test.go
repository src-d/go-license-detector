package processors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTML(t *testing.T) {
	text := `<body><h1>Title<h1/>
<p>some text</p>
&nbsp;&copy;&hellip;
<h2>Another title</h2>
<h3>And a third one.</h3>
<span>blah blah</span> <a href="wow://">LINK</a> text proceeds<br>
<script>here goes smth nasty<h4>even a title!</h4></script>hello
<style>you should put styles in &lt;head&gt;</style></body>`
	assert.Equal(t, `Title
some text
 ©…
Another title.
And a third one.
blah blah LINK text proceeds

hello
`, string(HTML([]byte(text))))
}
