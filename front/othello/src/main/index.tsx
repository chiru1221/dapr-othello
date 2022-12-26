import { useState, useReducer, useEffect } from 'react';
import { MainProps, gameAction,  playerState, playerAction, boardState, boardAction, Stone } from '../interface';
import { renderStone, countStone } from './stone';
import { strToArray, BoardColor } from './share';
import { PutableApi, CpApi, ReverseApi } from './api';

const Main = (props: MainProps) => {
    const initialBoardState = Array.from(
        Array(8), _ => Array(8).fill('n')
    ).map((values, index) => {
        const centerValues: string[] = values
        if (index === 3) {
            centerValues[3] = 'b'
            centerValues[4] = 'w'
        } else if (index === 4) {
            centerValues[3] = 'w'
            centerValues[4] = 'b'
        }
        return centerValues
    })
    const gameReducer = (state: [], action: gameAction) => {
        if (action.playerAction.player === 'cp') {
            // 1. request `/reverse`
            const reverseApi = ReverseApi(action.boardAction.stone, action.boardAction.i, action.boardAction.j, board.flat().join(''))
            // 2. request `/cp`
            reverseApi.then(reverse => {
                boardDispatch({i: -1, j: -1, stone: 'n', board: strToArray(reverse.squares)})
                const cpApi = CpApi(props.firstAttack? 'w': 'b', reverse.squares)
                cpApi.then(cp => {
                    // cp pass
                    if (parseInt(cp.x) === -1) {
                        const putableApi = PutableApi(props.firstAttack? 'b': 'w', reverse.squares)
                        putableApi.then(data => {
                            boardDispatch({
                                i: -1,
                                j: -1,
                                stone: 'n',
                                board: strToArray(data.squares),
                            })
                        })
                        playerDispatch({player: 'user'})
                        setPass(!pass)
                        return
                    }
                    // 3. boardDispatch
                    boardDispatch({i: parseInt(cp.x), j: parseInt(cp.y), stone: props.firstAttack? 'w': 'b'})
                    // 4. request `/reverse`
                    const newBoard = strToArray(reverse.squares)
                    newBoard[parseInt(cp.x)][parseInt(cp.y)] = props.firstAttack? 'w': 'b'
                    const reverseApi = ReverseApi(props.firstAttack? 'w': 'b', parseInt(cp.x), parseInt(cp.y), newBoard.flat().join(''))
                    reverseApi.then(reverse => {
                        boardDispatch({i: -1, j: -1, stone: 'n', board: strToArray(reverse.squares)})
                        // 5. request `/putable`
                        const putableApi = PutableApi(props.firstAttack? 'b': 'w', reverse.squares)
                        putableApi.then(data => {
                            boardDispatch({
                                i: -1,
                                j: -1,
                                stone: 'n',
                                board: strToArray(data.squares),
                            })
                            // user pass
                            if (!data.squares.includes('p')) {
                                playerDispatch({player: 'cp'})
                                setPass(!pass)
                                return
                            }
                            // 6. player Dispatch
                            playerDispatch({player: 'user'})
                        })
                    })
                })
            })
        }
        if (action.playerAction.player === 'user') {
            // 1. request `/putable`
            const putableApi = PutableApi(props.firstAttack? 'b': 'w', board.flat().join(''))
            // 2. boardDispatch
            putableApi.then(data => {
                // user pass
                if (!data.squares.includes('p')) {
                    return
                }
                boardDispatch({
                    i: -1,
                    j: -1,
                    stone: 'n',
                    board: strToArray(data.squares),
                })
            })
        }
        return state
    }
    const boardReducer = (state: boardState, action: boardAction) => {
        if (action.board) return [...action.board]
        if (action.stone === 'n') alert("????")
        const newState = state
        newState[action.i][action.j] = action.stone
        return [...newState]
    }
    const playerReducer = (state: playerState, action: playerAction) => {
        return action.player
    }
    const [board, boardDispatch] = useReducer(boardReducer, initialBoardState)
    const [player, playerDispatch] = useReducer(playerReducer, props.firstAttack? 'user': 'cp')
    const [game, gameDispatch] = useReducer(gameReducer, [])
    const [pass, setPass] = useState<boolean>(false)
    useEffect(() => {
        gameDispatch({
            playerAction: {player: player},
            boardAction: {i: -1, j: -1, stone: 'n', board: board}
        })
    }, [pass])
    
    return (
        <div style={{
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'center',
            alignItems: 'center',
            height: '100vh',
            backgroundColor: 'whitesmoke',
        }}>
            <div style={{height: '25%'}}>
                {countStone(board)}
            </div>
            {
                <div style={{
                    display: 'flex',
                    flexDirection: 'column',
                    backgroundColor: BoardColor
                }}>
                {board.map(
                    (row, i) => {
                        return (
                            <div key={`board-${i}`} style={{
                                display: 'flex',
                                flexDirection: 'row',
                            }}>
                            {row.map((col, j) => <div key={`board-${i}-${j}`} style={{border: '1px solid #424242'}}>{
                                renderStone({i: i, j: j, stone: col as Stone}, boardDispatch, playerDispatch, gameDispatch,props.firstAttack)
                            }</div>)}
                            <br />
                            </div>
                        )
                    }
                )}
                </div>
            }
        </div>
    )
}

export default Main