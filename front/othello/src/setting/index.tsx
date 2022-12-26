import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormLabel from '@mui/material/FormLabel';
import Button from '@mui/material/Button';
import { SettingProps } from '../interface';

const Setting = (props: SettingProps) => {
    const onClick = () => props.setSetting(false)
    return (
        <div style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'center',
            overflow: 'scroll',
            height: 'inherit',
        }}>
            <div>
                <FormLabel color='success'>First Attack</FormLabel>
                <RadioGroup
                    defaultValue={'user'}
                    name="radio-buttons-group"
                    onChange={(_, value) => props.setFirstAttack(value === 'user')}
                >
                    <FormControlLabel value={'user'} control={<Radio color="success" />} label="User" />
                    <FormControlLabel value={'cp'} control={<Radio color="success" />} label="CP" />
                </RadioGroup>
            </div>
            <div style={{marginTop: '5%'}}>
                <Button variant="outlined" color="success" onClick={onClick}>
                    OK
                </Button>
            </div>
        </div>
    )
}

export default Setting
