import Vue from 'vue';
import Entry from './vue/entry.vue';

/* Styles */
require.context('./styles/', false, /.*/);

/* Images */
require.context('./img/', false, /.*/);

/* Fonts */
require.context('./fonts/', false, /.*/);


new Vue({
    el: '#main-blog',
    components: { Entry }
});