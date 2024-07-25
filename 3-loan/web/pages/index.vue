<template>
    <main class="container max-w-2xl mx-auto">
    </main>
</template>

<script>
export default {
    data: function () {
        return {
            cards: [],
        }
    },
    mounted: function() {
        this.getCards()
    },
    methods: {
        getCards: function() {
            fetch('http://localhost:8000/api/v1/cards')
                .then(response => response.json())
                .then(json => {
                    this.cards = json.data
                })
        },
        addCard: function() {},
        review: function() {
            this.$router.push('/learn')
        },
    }
};
</script>