// Converts snake_case to camelCase
export const snakeToCamel = (str: string): string => {
  const split = str.split('_')

  return split.reduce((acc, curr) => {
    if (curr) {
      return acc + curr[0].toUpperCase() + curr.slice(1)
    }
    return acc
  }, split.shift() ?? '')
}

// Converts camelCase to snake_case
export const camelToSnake = (str: string): string => {
  const split = str.match(/^[^A-Z]+|[A-Z][^A-Z]*/g)

  if (!split) {
    return str
  }

  return split.map((s) => s.toLowerCase()).join('_')
}

// Converts an object and it's nested object keys from snake_case to camelCase
export const nestedSnakeToCamel = (obj: any): any => {
  if (obj === null) {
    return obj
  }

  const result = new obj.constructor()

  switch (obj.constructor.name) {
    case 'Object':
      for (const key in obj) {
        result[snakeToCamel(key)] = nestedSnakeToCamel(obj[key])
      }
      return result
    case 'Array':
      for (let i = 0; i < obj.length; i++) {
        result[i] = nestedSnakeToCamel(obj[i])
      }
      return result
    default:
      return obj
  }
}

// Converts an object and it's nested object keys from camelCase to snake_case
export const nestedCamelToSnake = (obj: any): any => {
  if (obj === null) {
    return obj
  }

  const result = new obj.constructor()

  switch (obj.constructor.name) {
    case 'Object':
      for (const key in obj) {
        result[camelToSnake(key)] = nestedCamelToSnake(obj[key])
      }
      return result
    case 'Array':
      for (let i = 0; i < obj.length; i++) {
        result[i] = nestedCamelToSnake(obj[i])
      }
      return result
    default:
      return obj
  }
}
