<template>
  <section>
    <button v-on:click="getPost">Send</button>
    <span class="entry-date">{{ date }}</span>
    <h1 class="entry-title">{{ title }}</h1>
    <div class="entry-content" v-html="text"></div>
  </section>
</template>

<script>
export default {
  data: function() {
    return {
      date: '23. Juni 2017',
      title: 'Apfelstreuselkuchen',
      text: 'Hello World<br>Wie geht es dir?'
    }
  },
  methods: {
    getPost: function() {
      this.$http.get('http://localhost:8080').then(() => {
        console.log('SUCCESS');
      });
    }
  }
}
</script>

<style>
  .message {
    color: blue;
  }
</style>