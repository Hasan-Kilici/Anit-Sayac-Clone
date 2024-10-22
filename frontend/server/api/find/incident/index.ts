import type { H3Event } from 'h3'

export default cachedEventHandler(async (event: H3Event) => {
  try {
    const { id } = getQuery(event)
    if (!id){
      throw createError({
        statusCode: 400,
        statusMessage: 'Bad Request'
      })
    }
    const url = `http://localhost:4000/api/find/incident?id=${id}` 

    const response = await $fetch(url)

    return response
  } catch (error) {
    throw createError({
      statusCode: 404,
      statusMessage: 'Not Found'
    })
  }
}, {
  maxAge: 60 * 60,
  getKey: (event: H3Event) => event.path
})
