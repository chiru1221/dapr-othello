export const BoardColor = 'darkgray'

export const strToArray = (str: string) => {
    const arr: string[][] = []
    for (let i = 0; i < 8; i++) {
        arr.push(Array.from(str.slice(8 * i, 8 * (i + 1))))
    }
    return arr
}

