$(document)
  .ready(function() {

    $('.filter.menu .item')
      .tab()
    ;

    $("#channel-list .item").each(
        function () { link=$(this).attr("link"); $(this).attr("onclick", "location.href='"+window.location.origin+link+"';") } )

  })
;

if (!window.location.origin)
    window.location.origin = window.location.protocol+"//"+window.location.host;
