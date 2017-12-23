import Vue from 'vue';
import VueResource from 'vue-resource';
import infiniteScroll from 'vue-infinite-scroll'


/* Vue Components */
import Entry from './vue/entry.vue';

/* Styles */
require.context('./styles/', false, /.*/);

/* Images */
require.context('./img/', false, /.*/);

/* Fonts */
require.context('./fonts/', false, /.*/);

/* Configure Vue */
Vue.use(VueResource);
Vue.use(infiniteScroll)

new Vue(
{
    /**
     * Use the div tag to dynamically load entries
     */
    el: '#main-blog',

    /**
     * Keep count of how many entries are displayed
     */
    data:
    {
        count: 0
    },

    /**
     * An entry component has the correct html markup, style declarations
     * and methods declared to fetch the content and display it.
     * 
     * Uses VueResource  to fetch the data from a mongodb database
     * at localhost:27017 which is served by a Go service
     */
    components:
    { 
        Entry
    },

    /**
     * The infiniteScroll plugin uses this event method when the user
     * scrolled down completely 
     */
    methods:
    {
        loadEntry: function()
        {
            this.$data.count++;
        }
    }
});