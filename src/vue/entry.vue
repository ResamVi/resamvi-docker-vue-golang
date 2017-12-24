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

    Vue resource has to be used, and it has to be version 0.9.3 because otherwise
    POST does not seem to work (empty request body)
    https://github.com/pagekit/vue-resource/issues/543#issuecomment-274419895
-->
<script>
export default
{
    data: function()
    {
        /**
         * Get entry data from database.
         * Access the count variable from the parent (entire entry blog)
         * to get the next entry in line to load
         */
        console.log("NEXT INDEX: " + this.$parent.count);
        let connection = this.$http.post('http://localhost:8080/', this.$parent.count);
    
        /*
        * The response is a JSON object with 
        * Number, Text, Date, Title properties (see server.go)
        */ 
        let onSuccess = function(response)
        {
            var entry = JSON.parse(response.body);
            
            this.$data.entry.Date = entry.Date;
            this.$data.entry.Title = entry.Title;
            this.$data.entry.Text = entry.Text;
            this.$data.entry.Number = entry.Number;
            //console.log(this.$data.entry);
            console.log(JSON.parse(response.body));
            //this.$data.entry = response.body;
            //console.log(this.$parent.count);
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
            entry : {Date: "Loading", Title:"Loading", Text:"Loading"},
        }
    }
}
</script>

<style>
    .entry-date {
        font-family: 'Cinzel';
        color: #696969;
    }

    .entry-title {
        margin-top:0px;
    }
    
    .entry-content {
        font-family: 'Open Sans';
        margin-bottom: 10%;
        font-size: 1.5em;
    }
</style>