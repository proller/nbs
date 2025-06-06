#pragma once

#include "public.h"

#include <cloud/blockstore/libs/diagnostics/public.h>

#include <cloud/storage/core/libs/common/public.h>
#include <cloud/storage/core/libs/common/startable.h>

namespace NCloud::NBlockStore::NServer {

////////////////////////////////////////////////////////////////////////////////

struct IEndpointProxyServer: public IStartable
{
};

////////////////////////////////////////////////////////////////////////////////

struct TEndpointProxyServerConfig
{
    ui16 Port;
    ui16 SecurePort;
    TString RootCertsFile;
    TString KeyFile;
    TString CertFile;
    TString UnixSocketPath;
    bool Netlink;
    TString StoredEndpointsPath;
    TDuration NbdRequestTimeout;
    TDuration NbdReconnectDelay;
    std::optional<ui32> DebugRestartEventsCount;

    TEndpointProxyServerConfig(
            ui16 port,
            ui16 securePort,
            TString rootCertsFile,
            TString keyFile,
            TString certFile,
            TString unixSocketPath,
            bool netlink,
            TString storedEndpointsPath,
            TDuration nbdRequestTimeout,
            TDuration nbdReconnectDelay,
            std::optional<ui32> debugRestartEventsCount)
        : Port(port)
        , SecurePort(securePort)
        , RootCertsFile(std::move(rootCertsFile))
        , KeyFile(std::move(keyFile))
        , CertFile(std::move(certFile))
        , UnixSocketPath(std::move(unixSocketPath))
        , Netlink(netlink)
        , StoredEndpointsPath(std::move(storedEndpointsPath))
        , NbdRequestTimeout(nbdRequestTimeout)
        , NbdReconnectDelay(nbdReconnectDelay)
        , DebugRestartEventsCount(debugRestartEventsCount)
    {
        Y_ENSURE(debugRestartEventsCount.value_or(1) > 0);
    }
};

////////////////////////////////////////////////////////////////////////////////

IEndpointProxyServerPtr CreateServer(
    TEndpointProxyServerConfig config,
    ITimerPtr timer,
    ISchedulerPtr scheduler,
    ILoggingServicePtr logging);

}   // namespace NCloud::NBlockStore::NServer
