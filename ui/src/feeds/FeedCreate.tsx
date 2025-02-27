import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Typography } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { usePageTitle } from '../hooks'
import FeedConfig from './FeedConfig'
import { FeedForm } from './Types'

const headers = {
  "Content-Type": "application/x-www-form-urlencoded",
}

export default withRouter(({ history }: RouteComponentProps) => {
  usePageTitle('new feed')
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  function handleBack() {
    setError(null)
    history.push('/feeds')
  }

  async function handleSave(form: FeedForm) {
    try {
      const { title, xmlUrl: url, tags } = form
      const res = await fetchAPI('/feeds', {title, url, tags}, {method: 'POST', headers})
      if (res.ok) {
        setError(null)
        const data = await res.json()
        showMessage(<Message variant="success"  message={`Feed ${data.title} (#${data.id}) created`} />)
        return history.push('/feeds')
      }
      const _err = await res.json()
      throw new Error(_err.detail || res.statusText)
    } catch (err) {
      setError(err)
    }
  }

  return (
    <>
      <Typography variant="h5" gutterBottom>New feed</Typography>
      { !!error && <Message message={error.message} variant="error" />}
      <FeedConfig onSave={handleSave} onCancel={handleBack} />
    </>
  )
})
