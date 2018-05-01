package processors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTML(t *testing.T) {
	text := `<body><h1>Title<h1/>
<p>some text</p>
<h2>Another title</h2>
<hr>
<h3>And a third one.</h3>
&nbsp; &gt; &lt; &amp; &quot; &apos; &cent; &pound; &yen; &euro; &copy; &reg;
&ldquo; &rdquo; &lsquo; &rsquo; &sbquo; &rbquo; &bdquo; &ndash; &mdash; &bull;
&hellip; &prime; &lsaquo; &rsaquo; &trade; &minus; &raquo; &laquo; &deg; &sect;
&iexcl;
<span>blah blah</span> <a href="wow://">LINK</a> text proceeds<br>
<script>here goes smth nasty<h4>even a title!</h4></script>hello
<style>you should put styles in &lt;head&gt;</style>
<a href="http://foo">http://foo</body>`
	assert.Equal(t, `Title
some text
Another title.
---
And a third one.
  > < & " ' ¢ £ ¥ € © ®
“ ” ‘ ’ ‚ " „ – — •
… ′ ‹ › ™ − » « ° §
¡
blah blah wow:// LINK text proceeds

hello

http://foo`, string(HTML([]byte(text))))
}
