import * as React from 'react'

import NextLink, { LinkProps as NextLinkProps } from 'next/link'
import MuiLink, { LinkProps as MuiLinkProps } from '@material-ui/core/Link'

type NextComposedProps = Omit<
  React.AnchorHTMLAttributes<HTMLAnchorElement>,
  'href'
> &
  NextLinkProps & { noDecoration?: boolean }

const NextComposed = React.forwardRef<HTMLAnchorElement, NextComposedProps>(
  (props, ref) => {
    const {
      as,
      href,
      replace,
      passHref,
      prefetch,
      noDecoration,
      style,
      ...other
    } = props

    return (
      <NextLink
        href={href}
        prefetch={prefetch}
        as={as}
        replace={replace}
        passHref={passHref}
      >
        <a
          ref={ref}
          style={noDecoration ? { ...style, textDecoration: 'none' } : style}
          {...other}
        />
      </NextLink>
    )
  },
)

type LinkProps = {
  innerRef?: React.Ref<HTMLAnchorElement>
  naked?: boolean
} & NextComposedProps &
  Omit<MuiLinkProps, 'href'>

const LinkContent = ({ href, innerRef, naked, ...other }: LinkProps) => {
  if (naked) {
    return <NextComposed ref={innerRef} href={href} {...other} />
  }
  return (
    <MuiLink
      component={NextComposed}
      ref={innerRef}
      href={href as string}
      {...other}
    />
  )
}

/** Custom link to use Next.js and Material-UI at the same time */
const Link = React.forwardRef<HTMLAnchorElement, LinkProps>(
  (props: LinkProps, ref) => <LinkContent {...props} innerRef={ref} />,
)

export default Link
