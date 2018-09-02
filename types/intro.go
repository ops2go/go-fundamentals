<!DOCTYPE html>
<!-- saved from url=(0024)https://play.golang.org/ -->
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<title>The Go Playground</title>
		<link rel="stylesheet" href="./intro_files/style.css">
		
		<script type="text/javascript" async="" src="./intro_files/analytics.js"></script><script async="" src="./intro_files/js"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());
			gtag('config', 'UA-11222381-7');
			gtag('config', 'UA-49880327-6');
		</script>
		
		<script src="./intro_files/jquery.min.js"></script>
		<script src="./intro_files/jquery-linedtextarea.js"></script>
		<script src="./intro_files/playground.js"></script>
		<script src="./intro_files/playground-embed.js"></script>
		<script>
		$(document).ready(function() {
			playground({
				'codeEl':       '#code',
				'outputEl':     '#output',
				'runEl':        '#run, #embedRun',
				'fmtEl':        '#fmt',
				'fmtImportEl':  '#imports',
				
				'shareEl':      '#share',
				'shareURLEl':   '#shareURL',
				
				'enableHistory': true,
				'enableShortcuts': true,
				'enableVet': true
			});
			playgroundEmbed({
				'codeEl':       '#code',
				
				'shareEl':      '#share',
				
				'embedEl':      '#embed',
				'embedLabelEl': '#embedLabel',
				'embedHTMLEl':  '#shareURL'
			});
			$('#code').linedtextarea();
			
			$('#code').attr('wrap', 'off');
			var about = $('#about');
			about.click(function(e) {
				if ($(e.target).is('a')) {
					return;
				}
				about.hide();
			});
			$('#aboutButton').click(function() {
				if (about.is(':visible')) {
					about.hide();
					return;
				}
				about.show();
			})
			
			if (readCookie('playgroundImports') == 'true') {
				$('#imports').attr('checked','checked');
			}
			$('#imports').change(function() {
				createCookie('playgroundImports', $(this).is(':checked') ? 'true' : '');
			});
			
			
			$('#run').click(function() {
				gtag('event', 'click', {
					event_category: 'playground',
					event_label: 'run-button',
				});
			});
			$('#share').click(function() {
				gtag('event', 'click', {
					event_category: 'playground',
					event_label: 'share-button',
				});
			});
			
		});

		function createCookie(name, value) {
			document.cookie = name+"="+value+"; path=/";
		}

		function readCookie(name) {
			var nameEQ = name + "=";
			var ca = document.cookie.split(';');
			for(var i=0;i < ca.length;i++) {
				var c = ca[i];
				while (c.charAt(0)==' ') c = c.substring(1,c.length);
				if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
			}
			return null;
		}
		</script>
	</head>
	<body itemscope="" itemtype="http://schema.org/CreativeWork">
		<input type="button" value="Run" id="embedRun">
		<div id="banner">
			<div id="head" itemprop="name">The Go Playground</div>
			<div id="controls">
				<input type="button" value="Run" id="run">
				<input type="button" value="Format" id="fmt">
				<div id="importsBox">
					<label title="Rewrite imports on Format">
						<input type="checkbox" id="imports">
						Imports
					</label>
				</div>
				
				<input type="button" value="Share" id="share">
				<input type="text" id="shareURL" style="display: none;">
				<label id="embedLabel" style="display: none;">
					<input type="checkbox" id="embed">
					embed
				</label>
				
			</div>
			<div id="aboutControls">
				<input type="button" value="About" id="aboutButton">
			</div>
		</div>
		<div id="wrap">
			<div class="linedtextarea" style="height:100%; overflow:hidden"><div class="lines" style="width: 3%; margin-top: -1023px;"><div class="">1</div><div>2</div><div>3</div><div>4</div><div>5</div><div class="">6</div><div>7</div><div>8</div><div class="">9</div><div>10</div><div>11</div><div>12</div><div>13</div><div>14</div><div>15</div><div>16</div><div>17</div><div>18</div><div>19</div><div>20</div><div>21</div><div>22</div><div>23</div><div>24</div><div>25</div><div>26</div><div>27</div><div>28</div><div>29</div><div>30</div><div>31</div><div>32</div><div>33</div><div>34</div><div>35</div><div>36</div><div>37</div><div>38</div><div>39</div><div>40</div><div>41</div><div>42</div><div>43</div><div>44</div><div>45</div><div>46</div><div>47</div><div>48</div><div>49</div><div>50</div><div>51</div><div>52</div><div>53</div><div>54</div><div>55</div><div>56</div><div>57</div><div>58</div><div>59</div><div>60</div><div>61</div><div>62</div><div>63</div><div>64</div><div>65</div><div>66</div><div>67</div><div>68</div><div>69</div><div class="">70</div><div>71</div><div>72</div><div>73</div><div>74</div><div>75</div><div>76</div><div>77</div><div>78</div><div>79</div><div>80</div><div>81</div><div>82</div><div>83</div><div>84</div><div>85</div><div>86</div><div>87</div><div>88</div><div>89</div><div>90</div><div>91</div><div>92</div><div>93</div><div>94</div><div>95</div><div>96</div><div>97</div><div>98</div><div>99</div><div>100</div><div>101</div><div>102</div><div>103</div><div>104</div><div>105</div><div>106</div><div>107</div><div>108</div><div>109</div><div>110</div><div>111</div><div>112</div><div>113</div><div>114</div><div>115</div><div>116</div><div>117</div><div>118</div><div>119</div><div>120</div><div>121</div><div>122</div><div>123</div><div>124</div><div>125</div><div>126</div><div>127</div><div>128</div><div>129</div><div>130</div><div>131</div><div>132</div><div>133</div><div>134</div><div>135</div><div>136</div><div>137</div><div>138</div><div>139</div><div>140</div><div>141</div><div class="">142</div><div>143</div><div>144</div><div>145</div><div>146</div><div>147</div><div>148</div><div>149</div><div>150</div><div>151</div><div>152</div><div>153</div><div>154</div><div>155</div><div>156</div><div>157</div><div>158</div><div>159</div><div>160</div><div>161</div><div>162</div><div>163</div><div>164</div><div>165</div><div>166</div><div>167</div><div>168</div><div>169</div><div>170</div><div>171</div><div>172</div><div>173</div><div>174</div><div>175</div><div>176</div><div>177</div><div>178</div><div>179</div><div>180</div><div>181</div><div>182</div><div>183</div><div>184</div><div>185</div><div>186</div><div>187</div><div>188</div><div>189</div><div>190</div><div>191</div><div>192</div><div>193</div><div>194</div><div>195</div><div>196</div><div>197</div><div>198</div><div>199</div><div>200</div><div>201</div><div>202</div><div>203</div><div>204</div><div>205</div></div><textarea itemprop="description" id="code" name="code" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off" style="width: 97%;">package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
</textarea></div>
		</div>
		<div id="output"><pre><span class="stdout">Hello World
1
40   1.61803398875
0.0010000000000000009
6 + 4 = 10
6 - 4 = 2
6 * 4 = 24
6 / 4 = 1
6 % 4 = 2
</span><span class="system">
Program exited.</span></pre></div>
		<img itemprop="image" src="./intro_files/gopher.png" style="display:none">
		<div id="about">
<p><b>About the Playground</b></p>

<p>
The Go Playground is a web service that runs on
<a href="https://golang.org/">golang.org</a>'s servers.
The service receives a Go program, compiles, links, and
runs the program inside a sandbox, then returns the output.
</p>

<p>
If the program contains <a href="https://golang.org/pkg/testing">tests or examples</a>
and no main function, the service runs the tests.
Benchmarks will likely not be supported since the program runs in a sandboxed
environment with limited resources.
</p>

<p>
There are limitations to the programs that can be run in the playground:
</p>

<ul>

<li>
The playground can use most of the standard library, with some exceptions.
The only communication a playground program has to the outside world
is by writing to standard output and standard error.
</li>

<li>
In the playground the time begins at 2009-11-10 23:00:00 UTC
(determining the significance of this date is an exercise for the reader).
This makes it easier to cache programs by giving them deterministic output.
</li>

<li>
There are also limits on execution time and on CPU and memory usage.
</li>

</ul>

<p>
The article "<a href="https://blog.golang.org/playground" target="_blank" rel="noopener">Inside
the Go Playground</a>" describes how the playground is implemented.
The source code is available at <a href="https://go.googlesource.com/playground" target="_blank" rel="noopener">
https://go.googlesource.com/playground</a>.
</p>

<p>
The playground uses the latest stable release of Go.<br>
The current version is <a href="https://play.golang.org/p/1VcPUlPk_3">go1.10.3</a>.
</p>

<p>
The playground service is used by more than just the official Go project
(<a href="https://gobyexample.com/">Go by Example</a> is one other instance)
and we are happy for you to use it on your own site.
All we ask is that you
<a href="mailto:golang-dev@googlegroups.com">contact us first (note this is a public mailing list)</a>,
use a unique user agent in your requests (so we can identify you),
and that your service is of benefit to the Go community.
</p>

<p>
Any requests for content removal should be directed to
<a href="mailto:security@golang.org">security@golang.org</a>.
Please include the URL and the reason for the request.
</p>
		</div>
	

</body></html>