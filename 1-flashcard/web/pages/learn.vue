<template>
    <main class="container max-w-2xl mx-auto">
        <div class="h-full flex flex-col">
            <div class="p-6">
                <div class="flex items-center justify-between">
                    <h1 class="text-3xl font-poppins font-black">
                        Reviewing Cards
                    </h1>
                </div>
            </div>
            <div class="card shadow-xl">
                <div class="card-body">
                    <div id="front">
                        <h2 class="font-poppins font-bold text-xs uppercase">Front</h2>
                        <p class="font-poppins font-bold text-3xl">{{ currentFront }}</p>
                    </div>
                    <div class="divider" v-show="isOpen"></div>
                    <div id="back" v-show="isOpen">
                        <h2 class="font-poppins font-bold text-xs uppercase">Back</h2>
                        <p class="font-poppins font-bold text-3xl">{{ currentBack }}</p>
                    </div>
                </div>
            </div>
            <div v-show="!isOpen" class="grid grid-cols-2 gap-4 pt-6">
                <button @click="fail" class="btn btn-md btn-error font-poppins shadow-xl">Fail</button>
                <button @click="pass" class="btn btn-md btn-success font-poppins shadow-xl">Pass</button>
            </div>
            <div v-show="isOpen" class="pt-6 flex flex-col">    
                <button @click="next" v-show="isOpen" class="btn btn-md btn-accent font-poppins shadow-xl">Next</button>
            </div>
        </div>
    </main>
</template>

<script>
export default {
    data: function () {
        return {
            cards: [],

            // App State
            currentId: 0,
            currentState: "CLOSE", // "OPEN" | "CLOSE"
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
        fail: function() {
            this.currentState = "OPEN"
        },
        pass: function() {
            this.currentState = "OPEN"
        },
        next: function() {
            this.currentId = this.currentId + 1
            this.currentState = "CLOSE"
        },
    },
    computed: {
        currentFront: function() {
            if (this.cards.length === 0) return ''
            return this.cards[this.currentId].front
        },
        currentBack: function() {
            if (this.cards.length === 0) return ''
            return this.cards[this.currentId].back
        },
        isOpen: function() {
            return this.currentState === "OPEN"
        },
    },
};
</script>