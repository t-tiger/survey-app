import React, { useContext, useState } from 'react'
import styled from 'styled-components'

import {
  Box,
  Button,
  ClickAwayListener,
  Divider,
  Grow,
  IconButton,
  MenuItem,
  MenuList,
  Paper,
  Popper,
  PopperPlacementType,
  Typography,
  useTheme,
} from '@material-ui/core'
import { ContactSupport, MoreVert, Person } from '@material-ui/icons'

import { Survey } from 'modules/survey/types'
import { formatDate } from 'utils/date'
import { totalVoteCount } from 'modules/survey/helpers'
import { useMessageCenter } from 'utils/messageCenter'
import { deleteSurvey } from 'modules/survey/api'

import AppContext from 'components/pages/AppContext'
import Link from 'components/atoms/Link'
import Router from 'next/router'
import ConfirmDialog from 'components/organisms/ConfirmDialog'

/** Material-UI's Popper child elements */
type PopperProps = {
  TransitionProps: {
    in: boolean
    onEnter: () => void
    onExited: () => void
  }
  placement: PopperPlacementType
}

type Props = {
  survey: Survey
  hideButton?: boolean
  onDelete: () => void
}

const SurveyItem: React.FC<Props> = ({
  survey,
  hideButton = false,
  onDelete,
}) => {
  const anchorRef = React.useRef<HTMLDivElement>(null)
  const [isOpenMenu, setOpenMenu] = useState(false)
  const [isOpenDeleteDialog, setOpenDeleteDialog] = useState(false)

  const { userId } = useContext(AppContext)
  const { showMessage } = useMessageCenter()
  const theme = useTheme()
  const editable = totalVoteCount(survey) === 0

  const handleClickEdit = () => {
    Router.push(`/surveys/${survey.id}/edit`)
  }
  const handleDelete = async () => {
    try {
      await deleteSurvey(survey.id)
      showMessage('success', 'Survey has been deleted successfully.')
      setOpenDeleteDialog(false)
      onDelete()
    } catch (e) {
      if (e.response?.data?.message) {
        showMessage('error', e.response.data.message)
      }
    }
  }

  return (
    <Container padding={2.5} boxShadow={2}>
      <Box display="flex" alignItems="start">
        <Typography variant="h6" style={{ lineHeight: 1.4, flex: 1 }}>
          {survey.title}
        </Typography>
        {survey.publisher.id === userId && (
          <div ref={anchorRef}>
            <IconButton
              style={{ margin: theme.spacing(-1.5) }}
              onClick={() => setOpenMenu(true)}
            >
              <MoreVert />
            </IconButton>
          </div>
        )}
      </Box>
      <Box mt={1} display="flex" color="#888">
        <Box display="flex">
          <Person fontSize="small" />
          <Typography
            variant="body2"
            style={{ marginLeft: theme.spacing(0.5) }}
          >
            {survey.publisher.name}
          </Typography>
        </Box>
        <Box display="flex" ml={1.5}>
          <ContactSupport fontSize="small" />
          <Typography
            variant="body2"
            style={{ marginLeft: theme.spacing(0.5) }}
          >
            {survey.questions.length} Questions
          </Typography>
        </Box>
      </Box>
      {!hideButton && (
        <Box mt={3}>
          {survey.publisher.id === userId ? (
            <Link href={`/surveys/${survey.id}/result`} noDecoration>
              <ActionButton
                variant="contained"
                color="secondary"
                disableElevation
              >
                Check Results
              </ActionButton>
            </Link>
          ) : (
            <Link href={`/surveys/${survey.id}/answer`} noDecoration>
              <ActionButton
                variant="contained"
                color="primary"
                disableElevation
              >
                Start Survey
              </ActionButton>
            </Link>
          )}
        </Box>
      )}
      <Box marginY={2}>
        <Divider />
      </Box>
      <Box display="flex" justifyContent="space-between" color="#888">
        <Typography variant="body2">
          {formatDate(new Date(survey.created_at))}
        </Typography>
        <Typography variant="body2">{totalVoteCount(survey)} Voted</Typography>
      </Box>
      <Popper open={isOpenMenu} anchorEl={anchorRef.current} transition>
        {({
          TransitionProps: { in: transIn, onEnter, onExited },
          placement,
        }: PopperProps) => (
          <Grow
            in={transIn}
            onEnter={onEnter}
            onExited={onExited}
            style={{
              transformOrigin:
                placement === 'bottom' ? 'center top' : 'center bottom',
            }}
          >
            <Paper>
              <ClickAwayListener onClickAway={() => setOpenMenu(false)}>
                <MenuList>
                  <MenuItem disabled={!editable} onClick={handleClickEdit}>
                    Edit Survey
                  </MenuItem>
                  <MenuItem onClick={() => setOpenDeleteDialog(true)}>
                    Delete Survey
                  </MenuItem>
                </MenuList>
              </ClickAwayListener>
            </Paper>
          </Grow>
        )}
      </Popper>
      <ConfirmDialog
        open={isOpenDeleteDialog}
        color="secondary"
        title="Delete survey"
        description="Are you sure to delete survey?"
        submitText="Delete"
        onClose={() => setOpenDeleteDialog(false)}
        onSubmit={handleDelete}
      />
    </Container>
  )
}

const Container = styled(Box)`
  border-radius: 8px;
  background: white;
`
const ActionButton = styled(Button)`
  width: 100%;
  border-radius: 20px;
`

export default SurveyItem
