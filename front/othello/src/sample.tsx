// Ref: https://github.com/roiban1344/react-tutorial-typescript-hooks/blob/main/src/index.tsx
import { useState } from 'react'
import { Repeat } from 'typescript-tuple'
import './game.css'

type SquareState = 'O' | 'X' | null

type SquareProps = {
    value: SquareState
    onClick: () => void
}

const Square = (props: SquareProps) => (
    <button className='square' onClick={props.onClick}>
        {props.value}
    </button>
)

type BoardState = Repeat<SquareState, 9>

type BoardProps = {
    squares: BoardState
    onClick: (i: number) => void
}

const Board = (props: BoardProps) => {
    const renderSquare = (i: number) => (
        <Square value={props.squares[i]} onClick={() => props.onClick(i)} />
    )

    return (
        <div>
            <div className='board-row'>
                {renderSquare(0)}
                {renderSquare(1)}
                {renderSquare(2)}
            </div>
            <div className='board-row'>
                {renderSquare(3)}
                {renderSquare(4)}
                {renderSquare(5)}
            </div>
            <div className='board-row'>
                {renderSquare(6)}
                {renderSquare(7)}
                {renderSquare(8)}
            </div>
        </div>
    )
}

type Step = {
    readonly squares: BoardState
    readonly xIsNext: boolean
}

type GameState = {
    readonly history: readonly Step[]
    readonly stepNumber: number
}

const Game = () => {
    const [state, setState] = useState<GameState>({
        history: [
            {
                squares: [null, null, null, null, null, null, null, null, null],
                xIsNext: true,
            },
        ],
        stepNumber: 0,
    })

    const current = state.history[state.stepNumber]
    // const winner = calculateWinner(current.squares)
    // let status: string
    // if (winner) {
    //     status = `Winner: ${winner}`
    // } else {
    //     status = `Next player: ${current.xIsNext ? 'X' : 'O'}`
    // }

    const handleClick = (i: number) => {
        const next: Step = (({ squares, xIsNext }) => {
            const nextSquares = squares.slice() as BoardState
            nextSquares[i] = xIsNext ? 'X' : 'O'
            return {
                squares: nextSquares,
                xIsNext: !xIsNext,
            }
        })(current)

        setState(({ history, stepNumber }) => {
            const newHistory = history.slice(0, stepNumber + 1).concat(next)

            return {
                history: newHistory,
                stepNumber: newHistory.length - 1,
            }
        })

        console.log(state);
    }

    const jumpTo = (move: number) => {
        setState(prev => ({
            ...prev,
            stepNumber: move,
        }))
    }

    return (
        <div className='game'>
            <div className='board'>
                <Board squares={current.squares} onClick={handleClick} />
            </div>
        </div>
    )
}

export default Game;
