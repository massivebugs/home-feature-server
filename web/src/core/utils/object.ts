import { camelCase, isArray, isObject, transform } from 'lodash'

// https://stackoverflow.com/questions/59769649/recursively-convert-an-object-fields-from-snake-case-to-camelcase
export const camelize = (obj: any) =>
  transform(obj, (acc: any, value, key: string, target) => {
    const camelKey = isArray(target) ? key : camelCase(key)

    acc[camelKey] = isObject(value) ? camelize(value) : value
  })
