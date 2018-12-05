import React, { PureComponent } from 'react'
import { colors } from '../../../../constants'
import { Screen, Header, PublicKeyWithActions } from '../../../Library'
import { paddingVertical } from '../../../../styles'
import { extractPublicKeyFromId } from '../../../../helpers/contacts'

export default class MyQRCode extends PureComponent {
  static navigationOptions = ({ navigation }) => ({
    tabBarVisible: false,
    header: <Header navigation={navigation} title='My QR code' backBtn />,
  })

  share = () => {
    console.error('share: not implemented')
  }

  render () {
    const { id, displayName } = this.props.navigation.getParam('data')
    const myID = extractPublicKeyFromId(id)
    return (
      <Screen style={[{ backgroundColor: colors.white }, paddingVertical]}>
        <PublicKeyWithActions initialKey={myID} initialName={displayName} copyButton shareButton readOnly self mode={'qrcode'} />
      </Screen>
    )
  }
}
