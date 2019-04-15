// @flow
import * as React from 'react'
import * as Sb from '../../stories/storybook'
import * as Kb from '../../common-adapters'
import RolePicker, {FloatingRolePicker} from './index'

let rolePickerProps = toMerge => ({
  disabledRoles: {},
  onCancel: Sb.action('cancel'),
  onConfirm: Sb.action('confirm'),
  onSelectRole: role => {
    Sb.action('onSelectRole - not attached to state')(role)
  },
  presetRole: 'Owner',
  ...toMerge,
})

class StateWrapper extends React.Component<any, any> {
  state = {}
  constructor() {
    super()
    this.state = {
      onSelectRole: role => {
        Sb.action('onSelectRole')(role)
        this.setState({selectedRole: role})
      },
    }
  }

  render() {
    return this.props.storyFn()(this.state, s => this.setState(s))
  }
}

const load = () => {
  Sb.storiesOf('Teams/Role Picker', module)
    .addDecorator(storyFn => <StateWrapper storyFn={storyFn} />)
    .add('Picker', () => state => (
      <RolePicker
        {...rolePickerProps({
          selectedRole: 'Owner',
          ...state,
        })}
      />
    ))
    .add('Picker - Disabled Owners', () => state => (
      <RolePicker
        {...rolePickerProps({
          disabledRoles: {
            Owner: 'Non-Keybase users can not be added as owners.',
          },
          onLetIn: Sb.action('Let in'),
          ...state,
        })}
      />
    ))
    .add('Picker - Header text', () => state => (
      <RolePicker
        {...rolePickerProps({
          headerText: 'Add them as:',
          onCancel: Sb.action('cancel'),
          ...state,
        })}
      />
    ))
    .add('Picker - With Let In / Cancel', () => state => (
      <RolePicker
        {...rolePickerProps({
          headerText: 'Add them as:',
          onCancel: Sb.action('cancel'),
          onLetIn: Sb.action('Let in'),
          ...state,
        })}
      />
    ))
    .add('Picker as popup dropdown from button', () => (state, setState) => (
      <Kb.Box2 direction="vertical" alignItems={'center'} style={{height: 600, justifyContent: 'flex-end'}}>
        <FloatingRolePicker
          position={'top center'}
          open={state.opened}
          {...rolePickerProps({
            headerText: 'Add them as:',
            onCancel: () => setState({opened: false}),
            onLetIn: Sb.action('Let in'),
            ...state,
          })}
        >
          <Kb.Button
            onClick={() => setState({opened: true})}
            disabled={state.opened}
            type="Primary"
            label="Add"
          />
        </FloatingRolePicker>
      </Kb.Box2>
    ))
}

export default load
