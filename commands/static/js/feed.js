var DagobahRouter = Backbone.Router.extend({
    routes: {
        // from: to
        "": "home",
        "search/*query": "search",
        "channel/*name": "channel",
        "post/*url": "post",
    },

    home: function() {
        this.app.collection.url = "/"
        this.app.collection.reset()
        // do something home
    },

    search: function(query) {
        this.app.collection.url = "/search/" + encodeURIComponent(query)
        this.app.collection.reset()
        // do some search
    },

    channel: function(name) {
        this.app.collection.url = "/channel/" + encodeURIComponent(name)
        this.app.collection.reset()
        // channel!
    },

    post: function(url) {
        this.app.collection.url = "/post/" + encodeURIComponent(url)
        this.app.collection.reset()
        // wheee
    },

    initialize: function(opts) {
        this.app = opts.app
    }
});

var Dagobah = Backbone.View.extend({
    el: 'body',

    events: {
        'click .js-navigate': 'navigate',
        'click .js-show-channels': 'showChannels',
        'click .js-show-latest': 'showItems',
        'click .js-navigate-channel': 'navigateChannel'
    },


    initialize: function() {
        _.bindAll(this, 'navigate')
        this.router = new DagobahRouter({ app: this});

    },

    showChannels: function(e) {
        $('#channel-list, .js-show-channels').addClass("active")
        $('#latest-list, .js-show-latest').removeClass("active")
    },

    showItems: function(e) {
        $('#channel-list, .js-show-channels').removeClass("active")
        $('#latest-list, .js-show-latest').addClass("active")
    },

    navigate: function(e) {
        debugger
        var uri = $(e.target).closest('a').attr('href');
        this.router.navigate(uri, {trigger: true});
        e.preventDefault();
    },

    navigateChannel: function(e) {
        var uri = $(e.target).closest('div').data('href');
        this.router.navigate(uri, {trigger: true});
        e.preventDefault();
    }
});

// 1 page worth of items
var Page = Backbone.Model.extend({
    url: function() {
       return this.collection.url + "?page="+ this.id
    }
})

var Pages = Backbone.Collection.extend({
    model: Page,

})

Dagobah.boot = function() {
    window.app = new Dagobah( { collection: new Pages } );
    Backbone.history.start({pushState: true});
}



//$(document)
  //.ready(function() {

    ////$('.filter.menu .item')
      ////.tab()
    ////;

    //$("#channel-list .item").each(
        //function () { link=$(this).attr("link"); $(this).attr("onclick", "location.href='"+window.location.origin+link+"';") } )

  //})
//;

//if (!window.location.origin)
    //window.location.origin = window.location.protocol+"//"+window.location.host;
