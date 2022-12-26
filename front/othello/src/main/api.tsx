export const PutableApi = (stone: string, squares: string) => {
    const putableApi = fetch('/putable', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'stone': stone,
            'squares': squares
        })
    })
    .then(res => res.json())
    return putableApi
}

export const CpApi = (stone: string, squares: string) => {
    const cpApi = fetch('/cp', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'stone': stone,
            'level': 3,
            'squares': squares
        })
    })
    .then(res => res.json())
    return cpApi
}

export const ReverseApi = (stone: string, x: number, y: number, squares: string) => {
    const reverseApi = fetch('/reverse', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'stone': stone,
            'x': x,
            'y': y,
            'squares': squares
        })
    })
    .then(res => res.json())
    return reverseApi
}
