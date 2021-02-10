// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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
import { Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'
import { merge } from 'lodash'

import Form from '@ttn-lw/components/form'
import SubmitBar from '@ttn-lw/components/submit-bar'
import SubmitButton from '@ttn-lw/components/submit-button'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import { RepositoryContext } from './context'
import ProgressHint from './hints/progress-hint'
import OtherHint from './hints/other-hint'
import Registration from './device-registration'
import Selection from './device-selection'
import Card from './device-card'
import reducer, {
  setBrand,
  setModel,
  setHwVersion,
  setFwVersion,
  setBand,
  setError,
  defaultState,
  selectVersion,
  selectBrand,
  selectModel,
  selectHwVersion,
  selectFwVersion,
  selectBand,
  selectError,
  hasAnySelectedOther,
  hasCompletedSelection,
} from './reducer'
import validationSchema, { initialValues } from './validation-schema'

import style from './repository.styl'

const m = defineMessages({
  selectDeviceTitle: '1. Select the end device',
  enterDataTitle: '2. Enter registration data',
  enterDataDescription:
    'Please choose an end device first to proceed with entering registration data',
  register: 'Register end device',
})

const stateToFormValues = state => ({
  ...initialValues,
  version_ids: {
    brand_id: selectBrand(state),
    model_id: selectModel(state),
    hardware_version: selectHwVersion(state),
    firmware_version: selectFwVersion(state),
    band_id: selectBand(state),
  },
})

const DeviceRepository = props => {
  const {
    appId,
    createDevice,
    createDeviceSuccess,
    getRegistrationTemplate,
    template,
    templateFetching,
    prefixes,
    mayEditKeys,
    jsConfig,
    asConfig,
    nsConfig,
    supportLink,
  } = props

  const asEnabled = asConfig.enabled
  const jsEnabled = jsConfig.enabled
  const nsEnabled = nsConfig.enabled
  const asUrl = asEnabled ? asConfig.base_url : undefined
  const jsUrl = jsEnabled ? jsConfig.base_url : undefined
  const nsUrl = nsEnabled ? nsConfig.base_url : undefined

  const [state, dispatch] = React.useReducer(reducer, defaultState)
  const version = selectVersion(state)
  const brand = selectBrand(state)
  const model = selectModel(state)
  const hardwareVersion = selectHwVersion(state)
  const firmwareVersion = selectFwVersion(state)
  const band = selectBand(state)
  const error = selectError(state)

  const versionRef = React.useRef(version)
  const formRef = React.useRef(null)
  const templateRef = React.useRef(template)

  const handleBrandChange = React.useCallback(({ value }) => dispatch(setBrand(value)), [])
  const handleModelChange = React.useCallback(({ value }) => dispatch(setModel(value)), [])
  const handleHwVersionChange = React.useCallback(({ value }) => dispatch(setHwVersion(value)), [])
  const handleFwVersionChange = React.useCallback(({ value }) => dispatch(setFwVersion(value)), [])
  const handleBandChange = React.useCallback(({ value }) => dispatch(setBand(value)), [])
  const handleSetError = React.useCallback(error => dispatch(setError(error)), [])

  const validationContext = React.useMemo(
    () => ({
      mayEditKeys,
      appId,
      asUrl,
      asEnabled,
      jsUrl,
      jsEnabled,
      nsUrl,
      nsEnabled,
    }),
    [appId, asEnabled, asUrl, jsEnabled, jsUrl, mayEditKeys, nsEnabled, nsUrl],
  )

  const handleSubmit = React.useCallback(
    async values => {
      try {
        const castedValues = validationSchema.cast(values, { context: validationContext })
        const { ids } = castedValues
        ids.application_ids = { application_id: appId }

        await createDevice(appId, castedValues)
        createDeviceSuccess(appId, ids.device_id)
      } catch (error) {
        handleSetError(error)
      }
    },
    [appId, createDevice, createDeviceSuccess, handleSetError, validationContext],
  )

  const hasSelectedOther = hasAnySelectedOther(state)
  const hasCompleted = hasCompletedSelection(state)
  const showProgressHint = !hasSelectedOther && !hasCompleted
  const showRegistrationForm = !hasSelectedOther && hasCompleted
  const showDeviceCard = !hasSelectedOther && hasCompleted && Boolean(template)
  const showOtherHint = hasSelectedOther

  const stateKey = React.useMemo(
    () => `${brand}-${model}-${hardwareVersion}-${firmwareVersion}-${band}`,
    [band, brand, firmwareVersion, hardwareVersion, model],
  )
  const stateKeyRef = React.useRef(stateKey)

  React.useEffect(() => {
    const version = selectVersion(state)
    const versionChanged = version !== versionRef.current
    // Reset version values if any have changed during end device selection.
    if (formRef.current && versionChanged) {
      formRef.current.setValues(stateToFormValues(state), false)
      versionRef.current = version
    }
  }, [getRegistrationTemplate, hasCompleted, state, template, version])

  React.useEffect(() => {
    const templateChanged = template !== templateRef.current
    // Merge the new device template with other form values.
    if (formRef.current && templateChanged) {
      const { end_device } = template
      formRef.current.setValues(merge(stateToFormValues(state), end_device), false)
      templateRef.current = template
    }
  }, [getRegistrationTemplate, hasCompleted, state, template, validationContext])

  React.useEffect(() => {
    // Fetch template after completing the selection step (select band, model, hw/fw versions and band).
    if (formRef.current && hasCompleted && stateKey !== stateKeyRef.current && !hasSelectedOther) {
      const {
        version_ids: { hardware_version, ...v },
      } = stateToFormValues(state)

      getRegistrationTemplate(appId, v)
      stateKeyRef.current = stateKey
    }
  }, [appId, getRegistrationTemplate, hasCompleted, hasSelectedOther, state, stateKey])

  return (
    <RepositoryContext.Provider value={state}>
      <Row>
        <Col>
          <Form
            formikRef={formRef}
            onSubmit={handleSubmit}
            initialValues={initialValues}
            validationSchema={validationSchema}
            validationContext={validationContext}
            error={error}
          >
            <Form.SubTitle title={m.selectDeviceTitle} />
            <Selection
              onBrandChange={handleBrandChange}
              onModelChange={handleModelChange}
              onHwVersionChange={handleHwVersionChange}
              onFwVersionChange={handleFwVersionChange}
              onBandChange={handleBandChange}
            />
            {showProgressHint && (
              <ProgressHint
                manualLinkPath={`/applications/${appId}/devices/add/manual`}
                supportLink={supportLink}
              />
            )}
            {showOtherHint && (
              <OtherHint
                manualLinkPath={`/applications/${appId}/devices/add/manual`}
                manualGuideDocsPath="/devices/adding-devices/"
              />
            )}
            {showDeviceCard && <Card brandId={brand} modelId={model} template={template} />}
            <hr className={style.hRule} />
            <Form.SubTitle title={m.enterDataTitle} />
            {showRegistrationForm ? (
              <Registration
                template={template}
                fetching={templateFetching}
                prefixes={prefixes}
                mayEditKeys={mayEditKeys}
              />
            ) : (
              <Message content={m.enterDataDescription} component="p" />
            )}
            <SubmitBar align="start">
              <Form.Submit
                message={m.register}
                component={SubmitButton}
                disabled={!showRegistrationForm}
              />
            </SubmitBar>
          </Form>
        </Col>
      </Row>
    </RepositoryContext.Provider>
  )
}

DeviceRepository.propTypes = {
  appId: PropTypes.string.isRequired,
  asConfig: PropTypes.stackComponent.isRequired,
  createDevice: PropTypes.func.isRequired,
  createDeviceSuccess: PropTypes.func.isRequired,
  getRegistrationTemplate: PropTypes.func.isRequired,
  jsConfig: PropTypes.stackComponent.isRequired,
  mayEditKeys: PropTypes.bool.isRequired,
  nsConfig: PropTypes.stackComponent.isRequired,
  prefixes: PropTypes.euiPrefixes.isRequired,
  supportLink: PropTypes.string,
  template: PropTypes.shape({
    end_device: PropTypes.shape({
      supports_join: PropTypes.bool,
      lorawan_version: PropTypes.string.isRequired,
    }),
  }),
  templateFetching: PropTypes.bool,
}

DeviceRepository.defaultProps = {
  template: undefined,
  templateFetching: false,
  supportLink: undefined,
}

export default DeviceRepository