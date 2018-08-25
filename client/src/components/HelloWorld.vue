<template>
    <div class="hello">
        <h1>{{ msg }}</h1>
        <input v-model="index"/>
        <button v-on:click="send()">send</button>
        <textarea>{{ listPi|json }}</textarea>
    </div>
</template>

<script>

    import { axiosProvider } from '../services';

    export default {
        name        : 'HelloWorld',
        props       : {
            msg: String,
        },
        data() {
            return {index: '', listPi: []}
        },
        filters     : {
            json: function (value) {
                return JSON.stringify(value)
            },
        },
        beforeCreate: function () {
            axiosProvider.request.get('/api/values/current').then(response => this.listPi = response.data)
        },
        methods     : {
            send: function () {
                axiosProvider.request.post('/api/values', {index: this.index})
            },
        },
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h3 {
        margin: 40px 0 0;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }
</style>
