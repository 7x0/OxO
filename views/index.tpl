<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="robots" content="index, nofollow">

  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-title" content="{{.appName}}">
  <meta name="apple-mobile-web-app-status-bar-style" content="black">

  <meta name="description" content="{{.appDescription}}">
  <meta name="keywords" content="">
  <meta property="og:title" content="{{.appName}}" />
  <meta property="og:image" content="static/images/apple-touch-icon.png" />
  <meta property="og:description" content="{{.appDescription}}" />
  <meta property="og:url" content="{{.appSite}}" />

  <title>{{.appName}}</title>

  <link rel="apple-touch-icon-precomposed" sizes="57x57" href="static/images/apple-touch-icon-57.png">
  <link rel="apple-touch-icon-precomposed" sizes="72x72" href="static/images/apple-touch-icon-72.png">
  <link rel="apple-touch-icon-precomposed" sizes="114x114" href="static/images/apple-touch-icon-114.png">
  <link rel="apple-touch-icon-precomposed" sizes="144x144" href="static/images/apple-touch-icon-144.png">

  <link rel="shortcut icon" href="static/images/favicon.png">

  <link rel="stylesheet" href="static/css/style.css">

</head>
<body>

  <a class="c" href="{{.appSite}}" tabindex="-1">{{.appName}}</a>

  <div id="wrap">
    <input id="url" type="text" placeholder="长地址" autocomplete="off">
    <div class="custom" style="display: none">
      <span>oxo.cat/</span>
      <input id="custom" type="text" placeholder="自定义一把" autocomplete="off">
      <span id="stat" style="display: none">不可用</span>
      <button id="meow" class="meow" type="submit">短短短</button>
    </div>
  </div>
  <div id="oxo" style="display: none">
    <span id="clip" value="">{{.appSite}}/<span id="res"></span></span>
    <button class="meow" data-clipboard-target="#clip">复制</button>
  </div>
  <span class="github">
    <iframe src="https://ghbtns.com/github-btn.html?user=7x0&repo=OxO&type=star&count=true" frameborder="0" scrolling="0" width="170px" height="20px"></iframe>
  </span>
  <span class="footer">
    Made with <span style="color: rgba(242, 101, 33, 1);">❤</span> in Amoy.
  </span>

  <script src="//cdn.bootcss.com/jquery/2.2.1/jquery.min.js"></script>
  <script src="//cdn.bootcss.com/clipboard.js/1.5.9/clipboard.min.js"></script>
  <script src="static/js/oxo.js"></script>
</body>
</html>
