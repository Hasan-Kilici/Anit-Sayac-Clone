import type { H3Event } from 'h3'

export default cachedEventHandler(async (event: H3Event) => {
  try {
    const { year } = getQuery(event)

    const url = year 
      ? `http://localhost:4000/api/list/incidents?year=${year}` 
      : 'http://localhost:4000/api/list/incidents'

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
