import React from 'react';


export interface SettingProps {
    setSetting: React.Dispatch<React.SetStateAction<boolean>>
    setFirstAttack: React.Dispatch<React.SetStateAction<boolean>>
}

export interface MainProps {
    firstAttack: boolean
}

export type gameState = {
    player: playerState,
    board: boardState,
}
export type gameAction = {
    playerAction: playerAction,
    boardAction: boardAction,
}
export type playerState = 'user' | 'cp' | 'master'
export type playerAction = {
    player: playerState
}
export type boardState = string[][]
export type boardAction = {
    i: number,
    j: number,
    stone: Stone,
    board?: boardState,
}
export type Stone = 'n' | 'b' | 'w' | 'p'
