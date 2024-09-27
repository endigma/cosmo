import { PlainMessage } from '@bufbuild/protobuf';
import { HandlerContext } from '@connectrpc/connect';
import { EnumStatusCode } from '@wundergraph/cosmo-connect/dist/common/common_pb';
import {
  GetSubgraphMetricsErrorRateRequest,
  GetSubgraphMetricsErrorRateResponse,
} from '@wundergraph/cosmo-connect/dist/platform/v1/platform_pb';
import { OrganizationRepository } from '../../repositories/OrganizationRepository.js';
import { SubgraphRepository } from '../../repositories/SubgraphRepository.js';
import { SubgraphMetricsRepository } from '../../repositories/analytics/SubgraphMetricsRepository.js';
import type { RouterOptions } from '../../routes.js';
import { enrichLogger, getLogger, handleError, validateDateRanges } from '../../util.js';

export function getSubgraphMetricsErrorRate(
  opts: RouterOptions,
  req: GetSubgraphMetricsErrorRateRequest,
  ctx: HandlerContext,
): Promise<PlainMessage<GetSubgraphMetricsErrorRateResponse>> {
  let logger = getLogger(ctx, opts.logger);

  return handleError<PlainMessage<GetSubgraphMetricsErrorRateResponse>>(ctx, logger, async () => {
    if (!opts.chClient) {
      return {
        response: {
          code: EnumStatusCode.ERR_ANALYTICS_DISABLED,
        },
        series: [],
      };
    }
    const authContext = await opts.authenticator.authenticate(ctx.requestHeader);
    logger = enrichLogger(ctx, logger, authContext);

    const subgraphMetricsRepo = new SubgraphMetricsRepository(logger, opts.chClient, opts.db);
    const subgraphRepo = new SubgraphRepository(logger, opts.db, authContext.organizationId);
    const orgRepo = new OrganizationRepository(logger, opts.db, opts.billingDefaultPlanId);

    const subgraph = await subgraphRepo.byName(req.subgraphName, req.namespace);
    if (!subgraph) {
      return {
        response: {
          code: EnumStatusCode.ERR_NOT_FOUND,
          details: `Subgraph '${req.subgraphName}' not found`,
        },
        series: [],
      };
    }

    const analyticsRetention = await orgRepo.getFeature({
      organizationId: authContext.organizationId,
      featureId: 'analytics-retention',
    });

    const { range, dateRange } = validateDateRanges({
      limit: analyticsRetention?.limit ?? 7,
      range: req.range,
      dateRange: req.dateRange,
    });

    const metrics = await subgraphMetricsRepo.getSubgraphErrorsView({
      range,
      dateRange,
      filters: req.filters,
      organizationId: authContext.organizationId,
      subgraphId: subgraph.id,
      subgraphLabels: subgraph.labels,
      namespaceId: subgraph.namespaceId,
    });

    return {
      response: {
        code: EnumStatusCode.OK,
      },
      series: metrics.errorRate.series,
      resolution: metrics.resolution,
    };
  });
}
