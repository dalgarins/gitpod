/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License.MIT.txt in the project root for license information.
 */

{
  grafanaDashboards+:: {
    'gitpod-sh-example-overview.json': (import 'dashboards/examples/overview.json'),
    'gitpod-sh-example-overview.json': (import 'dashboards/argocd/argocd.json'),
  },
}
