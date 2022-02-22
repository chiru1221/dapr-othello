// import { reverse } from 'dns/promises'
import { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import CircleIcon from '@mui/icons-material/Circle';
import Collapse from '@mui/material/Collapse';
import './game.css'

type BoardState = {
    readonly board: string
}

export type GameState = {
    readonly player: string,
    readonly bw: string,
    readonly cp: cpState
}

type cpState = {
    readonly level: number
}


// render using each player stone
const renderSquare = (square: string) => {
    let sx;
    if (square === 'b') {
        sx = {
            color: 'black',
        };
    } else if (square === 'w') {
        sx = {
            color: 'white',
        };
    } else if (square === 'p') {
        sx = {
            color: 'cyan',
        };
    }
    else {
        sx = {
            color: 'gray',
        };
    }
    
    return <div className='square'>
        <Collapse orientation="horizontal" in={true}>
            <CircleIcon sx={sx} />
        </Collapse>
    </div>;
}

// initialize board
// create board -> set stone at center
const initBoardArr = () => {
    let board = 'n'.repeat(27);
    board += 'bw';
    board += 'n'.repeat(6);
    board += 'wb';
    board += 'n'.repeat(27);
    return board;
}

const putAndReverse = async (
    i: number, j: number,
    game: GameState, board: BoardState
) => 
{
    // let newBoard = board.board.slice();
    // put
    let newBoard = board.board.slice(0, i*8+j);
    newBoard += game.bw;
    newBoard += board.board.slice(i*8+j+1)
    // reverse
    const url: string = '/reverse';
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'stone': game.bw,
            'x': i,
            'y': j,
            'squares': newBoard
        })
    });
    const r = await res.json();
    return r;
}

const putableSquare = async (
    game: GameState, board: BoardState
) => 
{
    // request for calculate putable square
    const url: string = '/putable';
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'stone': game.bw,
            'squares': board.board
        })
    });
    const r = await res.json();
    return r;
}

const cpAttakcs = async (
    game: GameState, board: BoardState
) => 
{
    // request for calculate putable square
    const url: string = "/cp";
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'stone': game.bw,
            'level': game.cp.level,
            'squares': board.board
        })
    });
    const r = await res.json();
    return r;
}

const Board = async (
    game: GameState, setGame: (gameState: GameState) => void,
    board: BoardState, setBoard: (boardState: BoardState) => void
) => {
    if (!board.board.includes('n')) {
        // end game
        const newUser = 'End';
        setGame({
            player: newUser,
            bw: game.bw,
            cp: game.cp,
        })
        return;
    } else if (!board.board.includes('b') || !board.board.includes('w')) {
        // end game
        const newUser = 'End';
        setGame({
            player: newUser,
            bw: game.bw,
            cp: game.cp,
        })
        return;
    }
    if (game.player === 'User') {
        console.log(game, board);
        // wait until user puts stone
        const newUser = 'Thinking';
        setGame({
            player: newUser,
            bw: game.bw,
            cp: game.cp
        });


        // print square that user can put
        const r = await putableSquare(game, board);
        setBoard({
            board: r.squares
        });
    } else if (game.player === 'CP') {
        console.log(game, board);
        // update game infromation
        const newUser = 'User';
        const newBW = (game.bw === 'b') ? 'w' : 'b';
        //
        // calculate score
        //
        setGame({
            player: newUser,
            bw: newBW,
            cp: game.cp
        });
        // cp attaks
        let r = await cpAttakcs(game, board);
        // when r.x == r.y == -1 -> cp pass
        if (r.x === -1) {
            r = {'squares': board.board};
        } else {
            r = await putAndReverse(r.x, r.y, game, board)
        }
        setBoard({
            board: r.squares
        });
    } else {
        console.log(game, board);
        if (board.board.includes('p')) {
            console.log('User thinking');
        } else {
            console.log('User pass');
            const newUser = 'CP';
            const newBW = (game.bw === 'b') ? 'w' : 'b';
            //
            // calculate score
            //

            setGame({
                player: newUser,
                bw: newBW,
                cp: game.cp
            });
            
            setBoard({
                board: board.board
            });
        }
    }
}

const renderBoard = (
    game: GameState, setGame: (gameState: GameState) => void,
    board: BoardState, setBoard: (boardState: BoardState) => void
) => {
    const squareClickFn = async (i: number, j: number) => {
        if (board.board[i*8+j] === 'p') {
            const newUser = 'CP';
            const newBW = (game.bw === 'b') ? 'w' : 'b';
            //
            // calculate score
            //
            setGame({
                player: newUser,
                bw: newBW,
                cp: game.cp
            });

            const r = await putAndReverse(i, j, game, board);
            setBoard({
                board: r.squares
            });
        }
    }

    const col: any[] = [];
    for (let i = 0; i < 8; i++) {
        const row: any[] = [];
        for (let j = 0; j < 8; j++) {
            row.push(
                <div onClick={() => squareClickFn(i, j)}>
                    {renderSquare(board.board[i*8+j])}
                </div>
            );
        }
        col.push(<div className='board-row'>{row}</div>)
    }
    return <div className='board'>{col}</div>;
}

const renderMeta = (game: GameState) => {
    return (
        <div className='meta'>
            <div className='meta-score'>
                Score
            </div>
            <div className='meta-player'>
                Player: {game.player}
            </div>
            <div className='meta-navigate'>

            </div>
        </div>
    )
}

const Game = () => {
    // user sets these settings at first
    const location: any = useLocation();
    let gameState: any;
    if (location.state) {
        gameState = location.state;
    } else {
        gameState = {
            player: 'User',
            bw: 'b',
            cp: {
                level: 1,
            }
        }
    }

    const [game, setGame] = useState<GameState>(
        gameState
    )
    const [board, setBoard] = useState<BoardState>(
        {
            board: initBoardArr()
        }
    );

    useEffect(() => {
        Board(game, setGame, board, setBoard);
    }, [board]);

    return (
        <div className='game'>
            { renderMeta(game) }
            { renderBoard(game, setGame, board, setBoard) }
        </div>
    );
}

export default Game;
