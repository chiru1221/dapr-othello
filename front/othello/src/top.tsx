import { useState } from 'react'
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormControl from '@mui/material/FormControl';
import FormLabel from '@mui/material/FormLabel';
import { useNavigate }from "react-router-dom";
import PlayArrowIcon from '@mui/icons-material/PlayArrow';
import { GameState } from './game';
import './top.css'


const Top = () => {
    const [select, setSelect] = useState<GameState>(
        {
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
    const navigate = useNavigate();
    
    const onStart = () => {
        navigate("/game", {state: select});
    }

    return (
        <div className='top'>
            <FormControl>
                <FormLabel>Attaks</FormLabel>
                <RadioGroup
                    aria-labelledby="radio-buttons-group-label"
                    defaultValue={'Pre-User'}
                    name="radio-buttons-group"
                    onChange={(_, value) => {
                        let newPlayer: string = value
                        setSelect({
                            player: newPlayer,
                            bw: 'b',
                            score: {
                                b: 2,
                                w: 2
                            },
                            cp: {
                                level: select.cp.level,
                                bw: (newPlayer === 'CP')? 'b': 'w'
                            }
                        })
                    }}
                >
                    <FormControlLabel value={'Pre-User'} control={<Radio />} label="First" />
                    <FormControlLabel value={'CP'} control={<Radio />} label="Second" />
                </RadioGroup>

                <FormLabel>Level</FormLabel>
                <RadioGroup
                    aria-labelledby="radio-buttons-group-label"
                    defaultValue={1}
                    name="radio-buttons-group"
                    onChange={(_, value) => {
                        let newCpLevel: number = +value;
                        setSelect({
                            player: select.player,
                            bw: 'b',
                            score: {
                                b: 2,
                                w: 2
                            },
                            cp: {
                                level: newCpLevel,
                                bw: select.cp.bw
                            }
                        })
                    }}
                >
                    <FormControlLabel value={1} control={<Radio />} label="Lv. 1" />
                    <FormControlLabel value={2} control={<Radio />} label="Lv. 2" />
                    <FormControlLabel value={3} control={<Radio />} label="Lv. 3" />
                </RadioGroup>
            </FormControl>
            
            <div onClick={onStart} className='game-start'>
                <PlayArrowIcon />
                <div>Start</div>
            </div>
        </div>
    )
}

export default Top;
