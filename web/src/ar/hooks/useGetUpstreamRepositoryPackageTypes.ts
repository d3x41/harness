/*
 * Copyright 2024 Harness, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import type { IconName } from '@harnessio/icons'

import { FeatureFlags } from '@ar/MFEAppTypes'
import type { StringsMap } from '@ar/frameworks/strings'
import { UpstreamProxyPackageType } from '@ar/pages/upstream-proxy-details/types'

import { useFeatureFlags } from './useFeatureFlag'

export interface UpstreamRepositoryPackageTypeListItem {
  label: keyof StringsMap
  value: UpstreamProxyPackageType
  icon: IconName
  disabled?: boolean
  tooltip?: string
  featureFlag?: FeatureFlags
}

export const useGetUpstreamRepositoryPackageTypes = (): UpstreamRepositoryPackageTypeListItem[] => {
  const featureFlags = useFeatureFlags()

  return UpstreamProxyPackageTypeList.map(repo => {
    if (repo.disabled && repo.featureFlag && featureFlags[repo.featureFlag]) {
      return {
        ...repo,
        disabled: false,
        tooltip: undefined
      }
    }
    return repo
  })
}

export const UpstreamProxyPackageTypeList: UpstreamRepositoryPackageTypeListItem[] = [
  {
    label: 'repositoryTypes.docker',
    value: UpstreamProxyPackageType.DOCKER,
    icon: 'docker-step'
  },
  {
    label: 'repositoryTypes.helm',
    value: UpstreamProxyPackageType.HELM,
    icon: 'service-helm'
  },
  {
    label: 'repositoryTypes.maven',
    value: UpstreamProxyPackageType.MAVEN,
    icon: 'maven-repository-type'
  },
  {
    label: 'repositoryTypes.npm',
    value: UpstreamProxyPackageType.NPM,
    icon: 'npm-repository-type',
    tooltip: 'Coming Soon!',
    featureFlag: FeatureFlags.HAR_NPM_PACKAGE_TYPE_ENABLED,
    disabled: true
  }
]
