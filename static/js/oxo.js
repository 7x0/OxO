function URLValidate(url) {
  var re = new RegExp(
        "^" +
          // protocol identifier
          "(?:(?:https?|ftp)://)" +
          // user:pass authentication
          "(?:\\S+(?::\\S*)?@)?" +
          "(?:" +
            // IP address exclusion
            // private & local networks
            "(?!(?:10|127)(?:\\.\\d{1,3}){3})" +
            "(?!(?:169\\.254|192\\.168)(?:\\.\\d{1,3}){2})" +
            "(?!172\\.(?:1[6-9]|2\\d|3[0-1])(?:\\.\\d{1,3}){2})" +
            // IP address dotted notation octets
            // excludes loopback network 0.0.0.0
            // excludes reserved space >= 224.0.0.0
            // excludes network & broacast addresses
            // (first & last IP address of each class)
            "(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])" +
            "(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}" +
            "(?:\\.(?:[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))" +
          "|" +
            // host name
            "(?:(?:[a-z\\u00a1-\\uffff0-9]-*)*[a-z\\u00a1-\\uffff0-9]+)" +
            // domain name
            "(?:\\.(?:[a-z\\u00a1-\\uffff0-9]-*)*[a-z\\u00a1-\\uffff0-9]+)*" +
            // TLD identifier
            "(?:\\.(?:[a-z\\u00a1-\\uffff]{2,}))" +
            // TLD may end with dot
            "\\.?" +
          ")" +
          // port number
          "(?::\\d{2,5})?" +
          // resource path
          "(?:[/?#]\\S*)?" +
        "$", "i"
      );
  return re.test(url);
}
$(document).ready(function(){
  $("#url").focus(function(){
    $(".custom").slideDown();
  });
  $("#url").blur(function(){
    var b = $("#meow");
    var s = $("#stat");
    if ($("#url").val().length == 0) {
      $(".custom").slideUp();
      s.hide();
      b.html("短短短");
    } else {
      if (!URLValidate($("#url").val())) {
        b.attr("disabled","true");
        s.show();
        s.html("不可用");
        b.html("Oh~ No!");
        $("#custom").attr("disabled","true");
      } else {
        s.hide();
        $("#custom").removeAttr("disabled");
        b.removeAttr("disabled");
        b.html("短短短");
      }
    }
  });
  // custom shorten URL check
  $("#custom").change(function(){
    var tag = $("#custom");
    var b = $("#meow");
    var s = $("#stat");
    if (tag.val().length != 0) {
      b.attr("disabled","true");
      $.post("/s/check", {
        custom:tag.val()
      },
      function(data, status){
        if (status) {
          if (!data) {
            s.show();
            s.html("不可用");
            b.html("Oh~ No!");
          } else {
            s.html("可用");
            b.removeAttr("disabled");
            b.html("短短短");
          }
        } else {
          s.show();
          s.html("炸毛了！");
        }
      });
    }
  });
  $("#meow").click(function(){
    if (URLValidate($("#url").val())) {
      var s = $("#stat");
      $.post("/s/gen", {
        url:$("#url").val(),
        custom:$("#custom").val()
      },
      function(data, status){
        if (status) {
          $("#wrap").hide();
          $("#oxo").show(700);
          $("#res").text(data);
        } else {
          s.show();
          s.html("炸毛了！");
        }
      });
    }
  });
});