<template>
    <main class="container max-w-2xl mx-auto">
        <div class="h-full flex flex-col">
            <div class="p-6">
                <div class="flex items-center justify-between">
                    <h1 class="text-3xl font-poppins font-black">
                        Flashcard
                    </h1>
                    <button class="btn btn-sm font-poppins">
                        Add Card
                    </button>
                </div>
            </div>
            <div class="card shadow-xl">
                <div class="card-body">
                    <h2 class="card-title font-poppins font-bold">Card List</h2>
                    <table class="table font-poppins">
                        <thead>
                            <tr>
                                <th></th>
                                <th>Front</th>
                                <th>Back</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(card, index) in cards">
                                <th>{{ index + 1 }}</th>
                                <td>{{ card.front }}</td>
                                <td>{{ card.back }}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
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
    }
};
</script>