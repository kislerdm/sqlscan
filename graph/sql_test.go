package graph_test

var inpt = `-- Check 'WHERE blst.client_id IS NULL' logic
-- Check with Christian costs net/gross
-- optimizer_standard_in insert
WITH
  bounds AS (
    SELECT client_id
          ,product_id
          ,unit_costs_net                               AS costs
          ,COALESCE(mcp_net, 0)                         AS mcp
          ,0.5                                          AS mcp_gap
          ,unit_costs_net > COALESCE(mcp_net, 0) + 0.5  AS bound_condition
    FROM facts.product_details
  )

INSERT INTO engine_pipeline.optimizer_standard_in (
    client_id
    ,product_id
    ,elasticity
    ,lower_bound
    ,upper_bound
    ,costs_total
)
SELECT e.client_id
      ,e.product_id
      ,e.elasticity
      ,CASE WHEN d.bound_condition THEN d.costs ELSE round((d.mcp + d.mcp_gap)::NUMERIC, 2) END  AS price_lower_bound
      ,CASE WHEN d.bound_condition THEN
            CASE
              WHEN d.mcp < 10 THEN round(1.5 * d.costs::NUMERIC, 2)
              WHEN d.mcp < 50 THEN round(1.3 * d.costs::NUMERIC, 2)
              WHEN d.mcp < 150 THEN round(1.2 * d.costs::NUMERIC, 2)
            ELSE round(1.15 * costs::NUMERIC, 2)
            END
      ELSE
          round(1.7 * (d.mcp::NUMERIC + 1.5), 2)
      END                                                                                       AS price_upper_bound
      ,d.costs                                                                                  AS costs_total --ask Christian where this to be NET, or GROSS
FROM engine_pipeline.elasticity_estimator_out AS e
INNER JOIN bounds AS d ON (d.client_id, d.product_id) = (e.client_id, e.product_id)
;
`
