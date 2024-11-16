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

// TODO: children don't seem to be converted well?
// TODO: make copy of object (this will solve the problem above)
// Converts an object and it's nested objects keys from snake to camel
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
