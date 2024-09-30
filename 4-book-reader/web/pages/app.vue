<template>
    <header class="py-5 px-5 flex justify-between border-b">
        <h1 class="text-2xl font-bold">Book Reader</h1>
        <label class="swap swap-rotate">
            <!-- this hidden checkbox controls the state -->
            <input type="checkbox" class="theme-controller" value="light" />
            <IconsSun class="swap-off" />
            <IconsMoon class="swap-on text-primary" />
        </label>
    </header>
    <main class="flex h-[90vh]">
        <div class="flex-1 p-5">
            <template v-for="word in content.split(' ')">
                <span class="border border-transparent hover:border-zinc-600 rounded-md p-[4px] cursor-pointer text-lg"
                    @click="handleWord">{{ word }}</span>&nbsp;
            </template>
        </div>
        <section class="flex-1 p-5 border-l">
            <div v-if="selectedWord">
                <div class="text-3xl font-bold mb-1">{{ selectedWordText }}</div>
                <div class="text-lg mb-6 dark:text-zinc-400">{{ definition.translation }}</div>
                <div v-if="definition.meaning" class="text-lg font-bold mb-1">Meaning</div>
                <p class="text-lg mb-4 dark:text-zinc-400">{{ definition.meaning }}</p>
                <button class="btn btn-sm btn-primary" @click="handleGet">
                    <span v-if="isGetting" class="loading loading-dots"></span>
                    Get via OpenAI
                </button>
            </div>
        </section>
    </main>
</template>

<script>
export default {
    data: function () {
        return {
            content: 'Akala niyo ba ang kapangyarihan ay nasa inyo sino ba kayo.',
            selectedWord: null,
            definition: {
                meaning: '',
                translation: '',
            },
            isGetting: false,
        }
    },
    methods: {
        handleWord: function (e) {
            if (this.selectedWord) {
                this.selectedWord.classList.remove('border-primary')
                this.selectedWord = null
                this.definition = {
                    meaning: '',
                    translation: '',
                }
            }
            this.selectedWord = e.currentTarget
            this.selectedWord.classList.add('border-primary')
        },
        handleGet: function () {
            this.isGetting = true
            $fetch('/api/definition', {
                method: 'POST',
                body: JSON.stringify({
                    word: this.selectedWordText
                })
            }).then(res => {
                this.definition = JSON.parse(res)
            }).finally(() => {
                this.isGetting = false
            })
        },
    },
    computed: {
        selectedWordText: function () {
            if (!this.selectedWord) {
                return ''
            }
            return this.selectedWord.innerText.toLowerCase()
                .replace(/[^\w\s]/gi, '')
        }
    }
}
</script>