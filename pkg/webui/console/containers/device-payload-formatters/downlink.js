// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import bind from 'autobind-decorator'
import { connect } from 'react-redux'

import PAYLOAD_FORMATTER_TYPES from '@console/constants/formatter-types'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { withBreadcrumb } from '@ttn-lw/components/breadcrumbs/context'
import toast from '@ttn-lw/components/toast'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import PayloadFormattersForm from '@console/components/payload-formatters-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import { updateDevice } from '@console/store/actions/devices'
import { attachPromise } from '@console/store/actions/lib'

import { selectSelectedApplicationId } from '@console/store/selectors/applications'
import {
  selectSelectedDeviceId,
  selectSelectedDeviceFormatters,
} from '@console/store/selectors/devices'

@connect(
  function(state) {
    const formatters = selectSelectedDeviceFormatters(state) || {}

    return {
      appId: selectSelectedApplicationId(state),
      devId: selectSelectedDeviceId(state),
      formatters,
    }
  },
  { updateDevice: attachPromise(updateDevice) },
)
@withBreadcrumb('device.single.payload-formatters.downlink', function(props) {
  const { appId, devId } = props

  return (
    <Breadcrumb
      path={`/applications/${appId}/devices/${devId}/payload-formatters/downlink`}
      content={sharedMessages.downlink}
    />
  )
})
class DevicePayloadFormatters extends React.PureComponent {
  static propTypes = {
    appId: PropTypes.string.isRequired,
    devId: PropTypes.string.isRequired,
    formatters: PropTypes.formatters.isRequired,
    updateDevice: PropTypes.func.isRequired,
  }

  @bind
  async onSubmit(values) {
    const { appId, devId, formatters, updateDevice } = this.props

    await updateDevice(appId, devId, {
      formatters: {
        down_formatter: values.type,
        down_formatter_parameter: values.parameter,
        up_formatter: formatters.up_formatter || PAYLOAD_FORMATTER_TYPES.NONE,
        up_formatter_parameter: formatters.up_formatter_parameter || '',
      },
    })

    toast({
      title: devId,
      message: sharedMessages.payloadFormattersUpdateSuccess,
      type: toast.types.SUCCESS,
    })
  }

  render() {
    const { formatters } = this.props

    return (
      <React.Fragment>
        <IntlHelmet title={sharedMessages.payloadFormattersDownlink} />
        <PayloadFormattersForm
          uplink={false}
          linked
          onSubmit={this.onSubmit}
          title={sharedMessages.payloadFormattersDownlink}
          initialType={formatters.down_formatter || PAYLOAD_FORMATTER_TYPES.NONE}
          initialParameter={formatters.down_formatter_parameter || ''}
        />
      </React.Fragment>
    )
  }
}

export default DevicePayloadFormatters
