import { OpenAI } from 'openai'


const client = new OpenAI({
    apiKey: process.env.OPENAI_API_KEY,
})


export default defineEventHandler(async (event) => {
    const { word } = await readBody(event)
    const definition = await _fetchDefinition(word)
    return definition
})


async function _fetchDefinition(word: str) {
    const response = await client.chat.completions.create({
        model: 'gpt-4o-mini',
        messages: [
            { role: 'system', content: 'Act as a Tagalog-English dictionary. Output response as JSON model {"meaning":"", "translation":""}. Only English. Do not output markdown.' },
            { role: 'user', content: word },
        ],
    })
    return response.choices[0].message.content
}
