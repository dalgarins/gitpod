/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { Container, ContainerModule } from "inversify";
import "mocha";
import * as chai from "chai";
import { PrebuildManager } from "./prebuild-manager";
import { TracedWorkspaceDB } from "@gitpod/gitpod-db/lib";
import { WorkspaceService } from "../workspace/workspace-service";
import { HostContextProvider } from "../auth/host-context-provider";
import { ConfigProvider } from "../workspace/config-provider";
import { Config } from "../config";
import { ProjectsService } from "../projects/projects-service";
import { IncrementalPrebuildsService } from "./incremental-prebuilds-service";
import { EntitlementService } from "../billing/entitlement-service";
import { CommitContext, Project, ProjectSettings, Repository, WorkspaceConfig } from "@gitpod/gitpod-protocol";

const expect = chai.expect;

const containerModule = new ContainerModule((bind) => {
    bind(PrebuildManager).toSelf().inSingletonScope();

    // #region: mocked dependencies of PrebuildManager
    bind(TracedWorkspaceDB).toConstantValue({} as any);
    bind(WorkspaceService).toConstantValue({} as any);
    bind(HostContextProvider).toConstantValue({} as any);
    bind(ConfigProvider).toConstantValue({} as any);
    bind(Config).toConstantValue({} as any);
    bind(ProjectsService).toConstantValue({} as any);
    bind(IncrementalPrebuildsService).toConstantValue({} as any);
    bind(EntitlementService).toConstantValue({} as any);
    // #endregion
});

const container = new Container();
container.load(containerModule);

const createPrebuildManager = () => container.get<PrebuildManager>(PrebuildManager);

describe("PrebuildManager", () => {
    const config: WorkspaceConfig = {
        _origin: "repo",
        tasks: [
            {
                name: "foo",
                init: "./foo",
            },
        ],
    };
    const context: CommitContext = {
        title: "",
        repository: {
            defaultBranch: "main",
        } as Repository,
        revision: "",
        ref: "main",
    };
    const settings: ProjectSettings = {
        enablePrebuilds: false,
    };
    const project = {
        settings,
    } as Project;

    function clone<T>(o: T, fn: (clone: T) => void) {
        const clone = JSON.parse(JSON.stringify(o));
        fn(clone);
        return clone;
    }

    const checkPrebuildPreconditionCases = [
        {
            title: "no-config",
            shouldRun: false,
            reason: "no-gitpod-config-in-repo",
            config: clone(config, (c) => (c._origin = undefined)),
            context,
            project,
        },
        {
            title: "no-tasks",
            shouldRun: false,
            reason: "no-tasks-in-gitpod-config",
            config: clone(config, (c) => (c.tasks = [])),
            context,
            project,
        },
        {
            title: "pre-existing-project/enable-by-default(1)",
            shouldRun: true,
            reason: "all-branches-selected",
            config,
            context,
            project: clone(project, (p) => (p.settings = undefined)),
        },
        {
            title: "pre-existing-project/enable-by-default(2)",
            shouldRun: true,
            reason: "all-branches-selected",
            config,
            context,
            project: clone(project, (p) => delete p.settings),
        },
        {
            title: "prebuilds-not-enabled",
            shouldRun: false,
            reason: "prebuilds-not-enabled",
            config,
            context,
            project: clone(project, (p) => (p.settings!.enablePrebuilds = false)),
        },
        {
            title: "default-branch-only/matched(1)",
            shouldRun: true,
            reason: "default-branch-matched",
            config,
            context,
            project: clone(
                project,
                (p) =>
                    (p.settings = {
                        enablePrebuilds: true,
                        prebuildDefaultBranchOnly: true,
                    }),
            ),
        },
        {
            title: "default-branch-only/matched(2)",
            shouldRun: true,
            reason: "default-branch-matched",
            config,
            context,
            project: clone(
                project,
                (p) =>
                    (p.settings = {
                        enablePrebuilds: true,
                        prebuildDefaultBranchOnly: undefined,
                    }),
            ),
        },
        {
            title: "default-branch-only/unmatched",
            shouldRun: false,
            reason: "default-branch-unmatched",
            config,
            context: clone(context, (c) => (c.ref = "feature-branch")),
            project: clone(
                project,
                (p) =>
                    (p.settings = {
                        enablePrebuilds: true,
                        prebuildDefaultBranchOnly: true,
                    }),
            ),
        },
        {
            title: "default-branch-only/default-branch-missing-in-context",
            shouldRun: false,
            reason: "default-branch-missing-in-commit-context",
            config,
            context: clone(context, (c) => delete c.repository.defaultBranch),
            project: clone(
                project,
                (p) =>
                    (p.settings = {
                        enablePrebuilds: true,
                        prebuildDefaultBranchOnly: true,
                    }),
            ),
        },
        {
            title: "all-branches/matched",
            shouldRun: true,
            reason: "all-branches-selected",
            config,
            context: clone(context, (c) => (c.ref = "feature-branch")),
            project: clone(
                project,
                (p) =>
                    (p.settings = {
                        enablePrebuilds: true,
                        prebuildDefaultBranchOnly: false,
                    }),
            ),
        },
    ];

    for (const { title, config, context, project, shouldRun, reason } of checkPrebuildPreconditionCases) {
        it(`checkPrebuildPrecondition/${title}`, async () => {
            const manager = createPrebuildManager();
            const precondition = manager.checkPrebuildPrecondition({ project, config, context });
            expect(precondition).to.deep.equal({ shouldRun, reason });
        });
    }
});