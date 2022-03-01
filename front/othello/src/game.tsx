// import { reverse } from 'dns/promises'
import { useState, useEffect } from 'react'
import { useLocation, Link } from 'react-router-dom'
import CircleIcon from '@mui/icons-material/Circle';
import RestartAltIcon from '@mui/icons-material/RestartAlt';
import HomeIcon from '@mui/icons-material/Home';
import Fade from '@mui/material/Fade';
import Collapse from '@mui/material/Collapse';
import './game.css'

type BoardState = {
    readonly board: string,
    readonly isReverse: boolean[][],
    readonly preBoard: string,
}

export type GameState = {
    readonly player: string,
    readonly bw: string,
    readonly score: scoreState
    readonly cp: cpState
}

type scoreState = {
    readonly b: number,
    readonly w: number
}

type cpState = {
    readonly level: number
    readonly bw: string
}


// render using each player stone
const renderSquare = (square: string, isIn: boolean) => {
    let sx = {'color': 'transparent', 'fontSize': '20px'};
    if (square === 'b') {
        sx.color = 'black';
    } else if (square === 'w') {
        sx.color = 'white';
    } else if (square === 'p') {
        sx.color = 'gray';
    }
    sx.fontSize = '4vh';

    const circle = (isIn)? (
        <Fade in={isIn} style={{transitionDelay: '0ms'}} {...(isIn ? { timeout: 1000 } : {})}>
            <CircleIcon sx={sx} />
        </Fade>
    ):(
        <CircleIcon sx={sx} />
    );
    
    return <div className='square'>
        {circle}
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

const countStone = async (squares: string) => {
    const blackScore = ((squares || '').match(/b/g) || []).length
    const whiteScore = ((squares || '').match(/w/g) || []).length
    return {
        b: blackScore,
        w: whiteScore,
    }
}

const diffBoard = async (preSquares: string, postSquares: string, putX: number, putY: number) => {
    console.log(putX, putY);
    const newIsIn = Array.from(Array(8), _ => Array(8).fill(false))
    for (let i = 0; i < 8; i++) {
        for (let j = 0; j < 8; j++) {
            if (preSquares[i*8+j] === 'p') {
                continue;
            }
            let isIn = (preSquares[i*8+j] === postSquares[i*8+j])? false: true;
            isIn = (putX === i && putY === j)? false: isIn;
            newIsIn[i][j] = isIn;
        }
    }
    return newIsIn;
}

const Board = async (
    game: GameState, setGame: (gameState: GameState) => void,
    board: BoardState, setBoard: (boardState: BoardState) => void
) => {
    // count stone
    const score = await countStone(board.board);
    console.log(score);
    // board before update
    const preBoard = board.board;

    // is game end?
    if (!board.board.includes('n')) {
        const newUser = 'End';
        setGame({
            player: newUser,
            bw: game.bw,
            score: score,
            cp: game.cp,
        })
        return;
    } else if (!board.board.includes('b') || !board.board.includes('w')) {
        const newUser = 'End';
        setGame({
            player: newUser,
            bw: game.bw,
            score: score,
            cp: game.cp,
        })
        return;
    }

    if (game.player === 'Pre-User') {
        console.log(game, board);
        // wait until user puts stone
        const newUser = 'User';
        // print square that user can put
        const r = await putableSquare(game, board);
        setTimeout(() => {
            setGame({
                player: newUser,
                bw: game.bw,
                score: score,
                cp: game.cp
            });

            setBoard({
                board: r.squares,
                isReverse: board.isReverse,
                preBoard: preBoard,
            });
        }, 1000);
    } else if (game.player === 'CP') {
        console.log(game, board);
        // new game infromation
        const newUser = 'Pre-User';
        const newBW = (game.bw === 'b') ? 'w' : 'b';
        // cp attaks
        let r = await cpAttakcs(game, board);
        //backup x,y
        const putX = r.x;
        const putY = r.y;
        // when r.x == r.y == -1 -> cp pass
        if (r.x === -1) {
            r = {'squares': board.board};
        } else {            
            r = await putAndReverse(r.x, r.y, game, board)
        }
        // new reverse
        const newIsReverse = await diffBoard(preBoard, r.squares, putX, putY);
        
        // update state
        setGame({
            player: newUser,
            bw: newBW,
            score: score,
            cp: game.cp
        });
        setTimeout(() => {
            setBoard({
                board: r.squares,
                isReverse: newIsReverse,
                preBoard: preBoard,
            });
        }, 1000);
    } else {
        console.log(game, board);
        if (board.board.includes('p')) {
            console.log('User thinking');
        } else {
            console.log('User pass');
            const newUser = 'CP';
            const newBW = (game.bw === 'b') ? 'w' : 'b';

            setGame({
                player: newUser,
                bw: newBW,
                score: score,
                cp: game.cp
            });
            setBoard({
                board: board.board,
                isReverse: board.isReverse,
                preBoard: board.preBoard
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
            const preBoard = board.board
            // new game information
            const newUser = 'CP';
            const newBW = (game.bw === 'b') ? 'w' : 'b';
            // user attack
            const r = await putAndReverse(i, j, game, board);
            // new revese
            const newIsReverse = await diffBoard(preBoard, r.squares, i, j);
            // update state
            setGame({
                player: newUser,
                bw: newBW,
                score: game.score,
                cp: game.cp
            });            
            setBoard({
                board: r.squares,
                isReverse: newIsReverse,
                preBoard: preBoard,
            });
        }
    }

    const col: any[] = [];
    for (let i = 0; i < 8; i++) {
        const row: any[] = [];
        for (let j = 0; j < 8; j++) {
            row.push(
                <div onClick={() => squareClickFn(i, j)}>
                    {renderSquare(board.board[i*8+j], board.isReverse[i][j])}
                </div>
            );
        }
        col.push(<div className='board-row'>{row}</div>)
    }
    return <div className='board'>{col}</div>;
}

const renderMeta = (game: GameState, board: BoardState) => {
    let player = (game.player === 'User')? 'Your Turn': 'CP';
    if (game.player === 'End') {
        // if same score -> winter is black
        const moreStone = (game.score.b >= game.score.w)? 'b': 'w';
        player = (moreStone === game.cp.bw)? 'You Lose': 'You Win';
    }
    return (
        <div className='meta'>
            <div className='meta-item'>
                CP Lv{game.cp.level}.
            </div>
            <div className='meta-item'>
                {player}
            </div>
            <div className='meta-item'>
                <CircleIcon sx={{fontSize: '5vh', color: 'black', marginRight: '1vw'}}/>
                <div>{game.score.b}</div>
                <CircleIcon sx={{fontSize: '5vh', color: 'white', margin: '0 1vw'}}/>
                <div>{game.score.w}</div>
            </div>
        </div>
    )
}

const renderNavigate = (game: GameState) => {
    return (
        <div className='navigate'>
            <Link to="/">
                <HomeIcon className='navigate-icon' sx={{fontSize: '8vh'}}/>
            </Link>
            <RestartAltIcon className='navigate-icon' sx={{fontSize: '8vh'}} onClick={() => window.location.reload()} />
        </div>
    )
}

const Game = () => {
    // user sets these settings at first
    const location: any = useLocation();

    const [game, setGame] = useState<GameState>(
        (location.state)? location.state: {
            player: 'Pre-User',
            bw: 'b',
            score: {
                b: 2,
                w: 2
            },
            cp: {
                level: 1,
                bw: 'w'
            }
        }
    )
    const [board, setBoard] = useState<BoardState>(
        {
            board: initBoardArr(),
            isReverse: Array.from(Array(8), _ => Array(8).fill(false)),
            preBoard: initBoardArr(),
        }
    );

    useEffect(() => {
        Board(game, setGame, board, setBoard);
    }, [board]);

    return (
        <div className='game'>
            { renderMeta(game, board) }
            <div className='board-wrap'>
                { renderBoard(game, setGame, board, setBoard) }
            </div>
            { renderNavigate(game) }
        </div>
    );
}

export default Game;
