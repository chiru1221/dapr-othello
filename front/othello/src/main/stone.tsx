import React from 'react';
import CircleIcon from '@mui/icons-material/Circle';
import CircleOutlinedIcon from '@mui/icons-material/CircleOutlined';
import Filter1Icon from '@mui/icons-material/Filter1';
import Filter2Icon from '@mui/icons-material/Filter2';
import Filter3Icon from '@mui/icons-material/Filter3';
import Filter4Icon from '@mui/icons-material/Filter4';
import Filter5Icon from '@mui/icons-material/Filter5';
import Filter6Icon from '@mui/icons-material/Filter6';
import Filter7Icon from '@mui/icons-material/Filter7';
import Filter8Icon from '@mui/icons-material/Filter8';
import Filter9Icon from '@mui/icons-material/Filter9';
import IconButton from '@mui/material/IconButton';
import { BoardColor } from './share';
import { boardAction, playerAction, gameAction, boardState } from '../interface';

export const renderStone = (
    action: boardAction,
    boardDispatch: React.Dispatch<boardAction>,
    playerDispatch: React.Dispatch<playerAction>,
    gameDispatch: React.Dispatch<gameAction>,
    firstAttack: boolean,
) => {
    if (action.stone === 'p') {
        const onClick = () => {
            const playerAction = {player: 'cp'} as playerAction
            const boardAction = {i: action.i, j: action.j, stone: firstAttack? 'b': 'w'} as boardAction
            playerDispatch(playerAction)
            boardDispatch(boardAction)
            gameDispatch({playerAction: playerAction, boardAction: boardAction})
        }
        return (
            <IconButton
                onClick={onClick}
            >
                <CircleOutlinedIcon />
            </IconButton>
        )
    } else if (action.stone === 'n') {
        return (
            <IconButton
                disabled={true}
            >
                <CircleIcon style={{color: BoardColor}}/>
            </IconButton>
        )
    }
    return (
        <IconButton
            disabled={true}
        >
            <CircleIcon style={{color: action.stone === 'w'? 'white': 'black'}} />
        </IconButton>
    )
}

export const countStone = (
    board: boardState
) => {
    const numberToIcon = [
        <Filter1Icon />,
        <Filter2Icon />,
        <Filter3Icon />,
        <Filter4Icon />,
        <Filter5Icon />,
        <Filter6Icon />,
        <Filter7Icon />,
        <Filter8Icon />,
        <Filter9Icon />,
    ]
    const renderWithIcon = (score: number) => {
        const strScore: string = score.toString()
        return (
            <div style={{display: 'flex', flexDirection: 'row'}}>
                {Array.from(strScore).map((key: string) => numberToIcon[parseInt(key) - 1])}
            </div>
        )
    }
    return (
        <div style={{
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'center',
            alignItems: 'center',
            width: '330px',
            backgroundColor: BoardColor,
            height: 'inherit',
            overflow: 'auto',
        }}>
            <CircleIcon style={{color: 'black', margin: '0 10%'}} />
            {renderWithIcon(board.flat().filter(stone => stone === 'b').length)}
            <CircleIcon style={{color: 'white', margin: '0 10%'}} />
            {renderWithIcon(board.flat().filter(stone => stone === 'w').length)}
        </div>
    )
}