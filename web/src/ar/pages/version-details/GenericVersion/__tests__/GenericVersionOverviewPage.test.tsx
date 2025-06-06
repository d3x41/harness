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

import React from 'react'
import userEvent from '@testing-library/user-event'
import { getByTestId, getByText, queryByTestId, render, waitFor } from '@testing-library/react'
import {
  type ArtifactDetail,
  type GenericArtifactDetailConfig,
  useGetArtifactDetailsQuery
} from '@harnessio/react-har-service-client'

import '@ar/pages/version-details/VersionFactory'
import '@ar/pages/repository-details/RepositoryFactory'
import { DEFAULT_DATE_TIME_FORMAT } from '@ar/constants'
import { getReadableDateTime } from '@ar/common/dateUtils'
import ArTestWrapper from '@ar/utils/testUtils/ArTestWrapper'

import VersionDetailsPage from '../../VersionDetailsPage'
import { mockGenericArtifactDetails, mockGenericVersionList, mockGenericVersionSummary } from './__mockData__'

jest.mock('@harnessio/react-har-service-client', () => ({
  useGetArtifactVersionSummaryQuery: jest.fn().mockImplementation(() => ({
    data: { content: mockGenericVersionSummary },
    error: null,
    isLoading: false,
    refetch: jest.fn()
  })),
  getAllArtifactVersions: jest.fn().mockImplementation(
    () =>
      new Promise(success => {
        success({ content: mockGenericVersionList })
      })
  ),
  useGetArtifactDetailsQuery: jest.fn().mockImplementation(() => ({
    data: { content: mockGenericArtifactDetails },
    error: null,
    isLoading: false,
    refetch: jest.fn()
  }))
}))

describe('Verify generic version overview page', () => {
  beforeAll(() => {
    jest.clearAllMocks()
  })

  test('Should not render integration detail cards', async () => {
    const { container } = render(
      <ArTestWrapper
        path="/registries/:repositoryIdentifier/artifacts/:artifactIdentifier/versions/:versionIdentifier/:versionTab"
        pathParams={{
          repositoryIdentifier: '1',
          artifactIdentifier: '1',
          versionIdentifier: '1',
          versionTab: 'overview'
        }}>
        <VersionDetailsPage />
      </ArTestWrapper>
    )

    expect(queryByTestId(container, 'integration-cards')).not.toBeInTheDocument()
  })

  test('should render overview content without error', async () => {
    const { container } = render(
      <ArTestWrapper
        path="/registries/:repositoryIdentifier/artifacts/:artifactIdentifier/versions/:versionIdentifier/:versionTab"
        pathParams={{
          repositoryIdentifier: '1',
          artifactIdentifier: '1',
          versionIdentifier: '1',
          versionTab: 'overview'
        }}>
        <VersionDetailsPage />
      </ArTestWrapper>
    )

    const generalInformationCard = getByTestId(container, 'general-information-card')
    expect(generalInformationCard).toBeInTheDocument()

    const data = mockGenericArtifactDetails.data as ArtifactDetail & GenericArtifactDetailConfig
    expect(generalInformationCard).toHaveTextContent('versionDetails.overview.generalInformation.title')

    expect(generalInformationCard).toHaveTextContent('versionDetails.overview.generalInformation.name' + data.name)
    expect(generalInformationCard).toHaveTextContent(
      'versionDetails.overview.generalInformation.version' + data.version
    )
    expect(generalInformationCard).toHaveTextContent(
      'versionDetails.overview.generalInformation.packageType' + 'packageTypes.genericPackage'
    )
    expect(generalInformationCard).toHaveTextContent(
      'versionDetails.overview.generalInformation.uploadedBy' +
        getReadableDateTime(Number(data.modifiedAt), DEFAULT_DATE_TIME_FORMAT)
    )
    expect(generalInformationCard).toHaveTextContent(
      'versionDetails.overview.generalInformation.description' + data.description
    )
  })

  test('should show error message if failed to load overview data', async () => {
    const refetchFn = jest.fn()
    ;(useGetArtifactDetailsQuery as jest.Mock).mockImplementation(() => ({
      data: null,
      error: { message: 'Failed to fetch artifact details' },
      isLoading: false,
      refetch: refetchFn
    }))
    const { container } = render(
      <ArTestWrapper
        path="/registries/:repositoryIdentifier/artifacts/:artifactIdentifier/versions/:versionIdentifier/:versionTab"
        pathParams={{
          repositoryIdentifier: '1',
          artifactIdentifier: '1',
          versionIdentifier: '1',
          versionTab: 'overview'
        }}>
        <VersionDetailsPage />
      </ArTestWrapper>
    )
    expect(getByText(container, 'Failed to fetch artifact details')).toBeInTheDocument()
    const retryBtn = container.querySelector('button[aria-label=Retry]')
    expect(retryBtn).toBeInTheDocument()
    await userEvent.click(retryBtn!)
    await waitFor(() => {
      expect(refetchFn).toHaveBeenCalled()
    })
  })
})
