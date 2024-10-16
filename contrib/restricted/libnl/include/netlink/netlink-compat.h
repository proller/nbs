/* SPDX-License-Identifier: LGPL-2.1-only */
/*
 * Copyright (c) 2003-2006 Thomas Graf <tgraf@suug.ch>
 */

#ifndef NETLINK_COMPAT_H_
#define NETLINK_COMPAT_H_

#ifdef __cplusplus
extern "C" {
#endif

#if !defined _LINUX_SOCKET_H && !defined _BITS_SOCKADDR_H
typedef unsigned short  sa_family_t;
#endif

#ifndef IFNAMSIZ
/** Maximum length of a interface name */
#define IFNAMSIZ 16
#endif

/* patch 2.4.x if_arp */
#ifndef ARPHRD_INFINIBAND
#define ARPHRD_INFINIBAND 32
#endif

/* patch 2.4.x eth header file */
#ifndef ETH_P_MPLS_UC
#define ETH_P_MPLS_UC  0x8847
#endif

#ifndef ETH_P_MPLS_MC
#define ETH_P_MPLS_MC   0x8848
#endif

#ifndef  ETH_P_EDP2
#define ETH_P_EDP2      0x88A2
#endif

#ifndef ETH_P_HDLC
#define ETH_P_HDLC      0x0019
#endif

#ifndef AF_LLC
#define AF_LLC		26
#endif

#ifndef AF_MPLS
#define AF_MPLS		28
#endif

#ifdef __cplusplus
}
#endif

#endif
