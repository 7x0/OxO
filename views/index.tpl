<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="robots" content="index, nofollow">

  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-title" content="HEX to RGB">
  <meta name="apple-mobile-web-app-status-bar-style" content="black">

  <meta name="description" content="A simple tool for converting HEX values to RGB and vice versa.">
  <meta name="keywords" content="hex, hexadecimal, rgb, converter, hex to rgb, hexadecimal to rgb, hex to rgb converter, rgb to hex, rgb to hex converter, convert hex to rgb, convert rgb to hex, hex to rgb conversion, rgb to hex conversion, hex converter, rgb converter, simple hex to rgb converter">
  <meta property="og:title" content="HEX to RGB Converter" />
  <meta property="og:image" content="static/images/apple-touch-icon.png" />
  <meta property="og:description" content="A simple tool for converting HEX values to RGB and vice versa." />
  <meta property="og:url" content="http://hex.colorrrs.com" />

  <title>OxO</title>

  <link rel="apple-touch-icon-precomposed" sizes="57x57" href="static/images/apple-touch-icon-57.png">
  <link rel="apple-touch-icon-precomposed" sizes="72x72" href="static/images/apple-touch-icon-72.png">
  <link rel="apple-touch-icon-precomposed" sizes="114x114" href="static/images/apple-touch-icon-114.png">
  <link rel="apple-touch-icon-precomposed" sizes="144x144" href="static/images/apple-touch-icon-144.png">

  <link rel="shortcut icon" href="static/images/favicon.png">

  <link rel="stylesheet" href="static/css/style.css">

</head>
<body>

  <a class="c" href="http://oxo.cat" tabindex="-1">OxO</a>

  <div id="wrap">
    <input id="url" type="text" placeholder="长地址" autocomplete="off">
    <div class="custom" style="display: none">
      <div>oxo.cat/</div>
      <input id="custom" type="text" placeholder="自定义" autocomplete="off">
      <button class="meow" type="submit">短短短</button>
    </div>
  </div>

  <span class="footer">Made with <span style="color: rgba(242, 101, 33, 1);">❤</span> in Amoy, by <a href="http://2x.io">Neo</a>.</span>

  <script src="//cdn.bootcss.com/jquery/2.2.1/jquery.min.js"></script>
  <script type="text/javascript">
    $(document).ready(function(){
      $("#url").focus(function(){
        $(".custom").slideDown();
      });
      $("#url").blur(function(){
        if ($("#url").val().length == 0) {
          $(".custom").slideUp();
        }
      });
    });
  </script>
</body>
</html>
