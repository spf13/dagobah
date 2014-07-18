var DagobahRouter = Backbone.Router.extend({
    routes: {
        // from: to
        "": "home",
        "search/*query": "search",
        "channel/*key": "channel",
        "post/*url": "post",
    },

    home: function() {
        this.app.sidebar.changeContext("home", "/", null);
    },

    search: function(query) {
        this.app.sidebar.changeContext("search", "/search/" + encodeURIComponent(query), query);
    },

    channel: function(key) {
        this.app.sidebar.changeContext("channel", "/channel/" + encodeURIComponent(key), key);
    },

    post: function(url) {
        var model = this.app.collection.add({id: url}, {silent: true});
        model.fetch({dataType: 'html'});
    },

    initialize: function(opts) {
        this.app = opts.app
    }
});

var ItemView = Backbone.View.extend({
    render: function() {
        this.$el.html(this.model.get('content'));
        return this;
    }
});

var MainView = Backbone.View.extend({
    el: '#main',

    initialize: function(opts) {
        this.app = opts.app;

        this.listenTo(this.collection, 'add', this.addItem);
        this.listenTo(this.collection, 'reset', this.render);
    },

    render: function() {
        this.$el.html('');
    },

    addItem: function(model) {
        var view = new ItemView({model: model});
        this.$el.append(view.render().el);
    }
});

var SidebarView = Backbone.View.extend({
    el: '#sidebar',

    events: {
        'click .js-show-channels': 'showChannels',
        'click .js-show-latest': 'showItems',
        'click .js-navigate-channel': 'navigateChannel'
    },

    initialize: function(opts) {
        this.app = opts.app
        this.$header = this.$('#header');
        this.$list = this.$('#latest-list');
    },

    context: null,
    url: null,

    nextItems: function() {

    },

    prevItems: function() {

    },

    showChannels: function(e) {
        this.$('#channel-list, .js-show-channels').addClass("active");
        this.$('#latest-list, .js-show-latest').removeClass("active");
    },

    showItems: function(e) {
        this.$('#channel-list, .js-show-channels').removeClass("active");
        this.$('#latest-list, .js-show-latest').addClass("active");
    },

    changeContext: function(context, url, token) {
        this.context = context
        this.url = url
        this.token = token

        this.updateHeader()
        this.showItems()

        var $list = this.$list
        $.ajax(this.url, {
            success: function(content){
                $list.html(content)
            }
        })
    },
    updateHeader: function() {
        this.$header.show()
        switch (this.context) {
            case 'channel':
                var label = this.$('#channel-list .item[data-key="'+this.token+'"] .name')
                this.$header.text(label.text())
                break;
            case 'search':

                break;
            case 'home':
                this.$header.hide()
                break;
        }
    },

    navigateChannel: function(e) {
        var uri = $(e.target).closest('div').data('href');
        this.app.router.navigate(uri, {trigger: true});
        e.preventDefault();
    }
})

var Dagobah = Backbone.View.extend({
    el: 'body',

    events: {
        //'click .js-navigate': 'navigate',
    },


    initialize: function() {
        _.bindAll(this, 'navigate');
        this.main = new MainView({app:this, collection: this.collection});
        this.sidebar = new SidebarView({ app: this });
        this.router = new DagobahRouter({ app: this });

    },

    navigate: function(e) {
        var uri = $(e.target).closest('a').attr('href');
        this.router.navigate(uri, {trigger: true});
        e.preventDefault();
    }
});

// 1 page worth of items
var Page = Backbone.Model.extend({
    url: function() {
       return this.collection.url + "?p="+ this.id
    }
})

var Pages = Backbone.Collection.extend({
    model: Page,

})

var Item = Backbone.Model.extend({
    url: function() {
        return '/post/' + encodeURIComponent(this.id);
    },
    parse: function(resp, opts) {
        return {content: resp};
    }
})

var Items = Backbone.Collection.extend({
    model: Item,

})

Dagobah.boot = function() {
    window.app = new Dagobah( { collection: new Items } );
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
