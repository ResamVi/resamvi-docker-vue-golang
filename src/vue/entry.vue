<!-- 
    An Entry consists of 
    - a date string in format DD.MONTHNAME.YYYY
    - a title string
    - the HTML body string

    "entry" is the json object with Date, Title, Text, Number properties.
    Names are case sensitive
-->
<template>
  <section>
    <span class="entry-date">{{ entry.Date }}</span>
    <h1 class="entry-title">{{ entry.Title }}</h1>
    <div class="entry-content" v-html="entry.Text"></div>
  </section>
</template>

<!--
    The data object consists of a single "entry" object
    of which the data is fetched from the database
    and stored in the appropriate properties "Date", "Title", "Text", "Number"
-->
<script>
export default
{
    data: function()
    {
        /**
         * Get entry data from database
         */
        let entry = {};
        let connection = this.$http.get('http://localhost:8080/')
    
        /*
        * The response is a JSON object with 
        * Number, Text, Date, Title properties (see server.go)
        */ 
        let onSuccess = function(response)
        {
            console.log(this.$parent.entries);
            this.$parent.entries.push(response.body);
            this.$data.entry = response.body;
        }

        /**
         * Error handling
         */
        let onFailure = function()
        {
            console.log('ERROR');
        }

        connection.then(onSuccess, onFailure);
        
        /**
         * While the entry is fetched from database (the code above is async)
         * display nothing.
         */
        return {
            entry : {Date: "", Title:"", Text:""}
        }
    }
}
</script>

<style>
  
</style>