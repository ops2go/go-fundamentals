<!DOCTYPE html>
<!-- saved from url=(0024)https://play.golang.org/ -->
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<title>The Go Playground</title>
		<link rel="stylesheet" href="./interfaces_files/style.css">
		
		<script type="text/javascript" async="" src="./interfaces_files/analytics.js"></script><script async="" src="./interfaces_files/js"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());
			gtag('config', 'UA-11222381-7');
			gtag('config', 'UA-49880327-6');
		</script>
		
		<script src="./interfaces_files/jquery.min.js"></script>
		<script src="./interfaces_files/jquery-linedtextarea.js"></script>
		<script src="./interfaces_files/playground.js"></script>
		<script src="./interfaces_files/playground-embed.js"></script>
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
			<div class="linedtextarea" style="height:100%; overflow:hidden"><div class="lines" style="width: 3%; margin-top: 0px;"><div class="">1</div><div>2</div><div class="">3</div><div>4</div><div>5</div><div class="">6</div><div>7</div><div class="">8</div><div class="">9</div><div>10</div><div>11</div><div>12</div><div>13</div><div>14</div><div>15</div><div>16</div><div>17</div><div>18</div><div>19</div><div>20</div><div>21</div><div>22</div><div>23</div><div>24</div><div class="">25</div><div>26</div><div>27</div><div class="">28</div><div>29</div><div>30</div><div>31</div><div>32</div><div>33</div><div class="">34</div><div>35</div><div>36</div><div>37</div><div class="">38</div><div>39</div><div class="">40</div><div>41</div><div>42</div><div>43</div><div>44</div><div>45</div><div>46</div><div>47</div><div>48</div><div>49</div><div>50</div><div>51</div><div>52</div><div>53</div><div>54</div><div>55</div><div>56</div><div>57</div><div>58</div><div>59</div><div>60</div><div>61</div><div>62</div><div>63</div><div>64</div><div>65</div><div>66</div><div class="">67</div><div>68</div><div>69</div><div class="">70</div><div>71</div><div>72</div><div>73</div><div>74</div><div>75</div><div>76</div><div>77</div><div class="">78</div><div>79</div><div>80</div><div>81</div><div>82</div><div>83</div><div class="">84</div><div>85</div><div>86</div><div>87</div><div>88</div><div>89</div><div>90</div><div>91</div><div class="">92</div><div>93</div><div>94</div><div>95</div><div>96</div><div>97</div><div>98</div><div>99</div><div>100</div><div class="">101</div><div>102</div><div>103</div><div>104</div><div>105</div><div>106</div><div>107</div><div>108</div><div>109</div><div class="">110</div><div>111</div><div>112</div><div>113</div><div>114</div><div>115</div><div>116</div><div>117</div><div>118</div><div class="">119</div><div>120</div><div class="">121</div><div>122</div><div>123</div><div>124</div><div>125</div><div class="">126</div><div>127</div><div>128</div><div>129</div><div>130</div><div>131</div><div>132</div><div>133</div><div>134</div><div>135</div><div class="">136</div><div>137</div><div>138</div><div>139</div><div>140</div><div>141</div><div class="">142</div><div>143</div><div>144</div><div>145</div><div>146</div><div>147</div><div class="">148</div><div>149</div><div>150</div><div>151</div><div>152</div><div>153</div><div>154</div><div>155</div><div>156</div><div>157</div><div>158</div><div>159</div><div>160</div><div>161</div><div>162</div><div>163</div><div>164</div><div>165</div><div>166</div><div>167</div><div>168</div><div>169</div><div>170</div><div>171</div><div>172</div><div>173</div><div>174</div><div>175</div><div>176</div><div>177</div><div>178</div><div>179</div><div>180</div><div>181</div><div>182</div><div>183</div><div>184</div><div>185</div><div>186</div><div>187</div><div>188</div><div>189</div><div>190</div><div>191</div><div>192</div><div>193</div><div>194</div><div>195</div><div>196</div><div>197</div><div>198</div><div>199</div><div>200</div><div>201</div><div>202</div><div>203</div><div>204</div><div>205</div><div>206</div><div>207</div><div>208</div><div>209</div><div>210</div><div>211</div><div>212</div><div>213</div><div>214</div><div>215</div><div>216</div><div>217</div><div>218</div><div>219</div><div>220</div><div>221</div><div>222</div><div>223</div><div>224</div><div>225</div><div>226</div><div>227</div><div>228</div><div>229</div><div>230</div><div>231</div><div>232</div><div>233</div><div>234</div><div>235</div><div>236</div><div>237</div><div>238</div><div>239</div><div>240</div><div>241</div><div>242</div><div>243</div><div>244</div><div>245</div><div>246</div><div>247</div><div>248</div><div>249</div><div>250</div><div>251</div><div>252</div><div>253</div><div>254</div><div>255</div><div>256</div><div>257</div><div>258</div><div>259</div><div>260</div><div>261</div><div>262</div><div>263</div><div>264</div><div>265</div><div>266</div><div>267</div><div>268</div><div>269</div><div>270</div><div>271</div><div>272</div><div>273</div><div>274</div><div>275</div><div>276</div><div>277</div><div>278</div><div>279</div><div>280</div><div>281</div><div>282</div><div>283</div><div>284</div><div>285</div><div>286</div><div>287</div><div>288</div><div>289</div><div>290</div><div>291</div><div>292</div><div>293</div><div>294</div><div>295</div><div>296</div><div>297</div><div class="">298</div><div>299</div><div>300</div><div>301</div><div>302</div><div>303</div><div>304</div><div>305</div><div>306</div><div>307</div><div>308</div><div>309</div><div>310</div><div>311</div><div>312</div><div>313</div><div>314</div><div>315</div><div>316</div><div>317</div><div>318</div><div>319</div><div>320</div><div>321</div><div>322</div><div>323</div><div>324</div><div>325</div><div>326</div><div>327</div><div>328</div><div>329</div><div>330</div><div>331</div><div>332</div><div>333</div><div>334</div><div>335</div><div>336</div><div>337</div><div>338</div><div>339</div><div>340</div><div>341</div><div>342</div><div>343</div><div>344</div><div>345</div><div>346</div><div>347</div><div>348</div><div>349</div><div>350</div><div>351</div><div>352</div><div>353</div><div>354</div><div>355</div><div>356</div><div>357</div><div>358</div><div class="">359</div><div>360</div><div>361</div><div>362</div><div>363</div><div>364</div><div>365</div><div>366</div><div>367</div><div>368</div><div>369</div><div>370</div><div>371</div><div>372</div><div>373</div><div>374</div><div>375</div><div class="">376</div><div>377</div><div>378</div><div>379</div><div>380</div><div>381</div><div>382</div><div>383</div><div class="">384</div><div>385</div><div>386</div><div>387</div><div>388</div><div>389</div><div>390</div><div>391</div><div>392</div><div>393</div><div>394</div><div>395</div><div>396</div><div>397</div><div>398</div><div>399</div><div>400</div><div>401</div><div class="">402</div><div>403</div><div>404</div><div>405</div><div>406</div><div>407</div><div>408</div><div>409</div><div>410</div><div class="">411</div><div>412</div><div class="">413</div><div>414</div><div>415</div><div>416</div><div>417</div><div class="">418</div><div>419</div><div>420</div><div>421</div><div>422</div><div>423</div><div>424</div><div>425</div><div>426</div><div>427</div><div class="">428</div><div>429</div><div>430</div><div>431</div><div>432</div><div>433</div><div>434</div><div>435</div><div>436</div><div>437</div><div>438</div><div>439</div></div><textarea itemprop="description" id="code" name="code" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off" style="width: 97%;">package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
</textarea></div>
		</div>
		<div id="output"><pre><span class="stdout">Rectangle Area = 1000
Circle Area = 50.26548245743669
</span><span class="system">
Program exited.</span></pre></div>
		<img itemprop="image" src="./interfaces_files/gopher.png" style="display:none">
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