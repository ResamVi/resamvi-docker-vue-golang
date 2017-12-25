import Vue from 'vue';
import infiniteScroll from 'vue-infinite-scroll';

/* Vue Components */
import Entry from './vue/entry.vue';

/* Styles */
require.context('./styles/', false, /.*/);

/* Images */
require.context('./img/', true, /.*/);

/* Fonts */
require.context('./fonts/', false, /.*/);

/* Configure Vue */
Vue.use(infiniteScroll);

new Vue({
    /**
     * Use the div tag to dynamically load entries
     */
    el: '#main-blog',

    /**
     * Keep count of how many entries are displayed.
     * 
     * The <div> tag with v-infinite-scroll in index.html will trigger a loadEntry event
     * when the user scrolled down to the main blog.
     * 
     * Indexing of entries start at one. Thus the first increment in loadEntry will
     * put the count variable at 1.
     */
    data: {
        count: 0,
        busy: false
    },

    /**
     * Used in index.html to display an entry entity
     * 
     * An entry component has the correct html markup, style declarations
     * and methods declared to fetch the content and display it.
     * 
     * Uses VueResource  to fetch the data from a mongodb database
     * at localhost:27017 which is served by a Go service
     */
    components: {
        Entry
    },

    /**
     * The infiniteScroll plugin uses this event method when the user
     * scrolled down completely 
     */
    methods: {
        loadEntry: function () {
            this.$data.count++;
            this.$data.busy = true;

            setTimeout(() => {
                this.busy = false;
            }, 10);
        }
    }
});