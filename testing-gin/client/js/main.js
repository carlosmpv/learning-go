new Vue({
    el: "#app",
    data () {
        return {
            products: null
        }
    },
    mounted () {
        axios
            .get("/api/list")
            .then(response => (this.products = response.data))
    }
});