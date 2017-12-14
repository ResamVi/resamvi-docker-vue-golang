import Vue from 'vue';
import VueResource from 'vue-resource';

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

new Vue({
    el: '#main-blog',
    components: { Entry },
});