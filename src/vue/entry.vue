<template>
  <section>
    <span class="entry-date">{{ date }}</span>
    <h1 class="entry-title">{{ title }}</h1>
    <div class="entry-content"> {{ text }} </div>
  </section>
</template>

<script>
export default {
  data () {
    return {
      date: '23. Juni 2017',
      title: 'Apfelstreuselkuchen',
      text: 'Hello World'
    }
  }
}
</script>

<style>
  .message {
    color: blue;
  }
</style>