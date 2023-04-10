document.querySelector('#btnSuccess').removeAttribute('disabled');
$("#btnSuccess").attr("style", "cursor:pointer;");
$("#btnSuccess").click(function () {
  var url = $(this).attr("href");
  gtag_report_conversion(url);
});

function gtag_report_conversion(url) {
  var callback = function () {
    if (typeof (url) != 'undefined') {
      window.location = url;
    }
  };

  gtag('event', 'conversion', {
    'send_to': 'AW-579666125/yeCfCOnuoeEBEM2BtJQC',
    'event_callback': callback
  });
  return false;
}