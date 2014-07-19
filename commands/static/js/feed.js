$( document ).ready(function() {
    $sidebar = $('#sidebar');

    $('#sidebar .js-show-channels').click( function() {
        $('#sidebar > #channel-list, .js-show-channels').addClass('active');
        $('#sidebar > #latest-list, .js-show-latest').removeClass('active');
    })
    $('#sidebar .js-show-latest').click( function() {
        $('#sidebar > #channel-list, .js-show-channels').removeClass('active');
        $('#sidebar > #latest-list, .js-show-latest').addClass('active');
    })
    $('#sidebar > #channel-list .item').each( function () {
        link=$(this).data('href'); $(this).attr('onclick', "location.href='"+window.location.origin+link+"';")
    })
});

if (!window.location.origin)
    window.location.origin = window.location.protocol+"//"+window.location.host;
