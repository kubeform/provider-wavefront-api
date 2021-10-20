/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Dashboard struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DashboardSpec   `json:"spec,omitempty"`
	Status            DashboardStatus `json:"status,omitempty"`
}

type DashboardSpecParameterDetails struct {
	DefaultValue *string `json:"defaultValue" tf:"default_value"`
	// +optional
	DynamicFieldType *string `json:"dynamicFieldType,omitempty" tf:"dynamic_field_type"`
	HideFromView     *bool   `json:"hideFromView" tf:"hide_from_view"`
	Label            *string `json:"label" tf:"label"`
	Name             *string `json:"name" tf:"name"`
	ParameterType    *string `json:"parameterType" tf:"parameter_type"`
	// +optional
	QueryValue *string `json:"queryValue,omitempty" tf:"query_value"`
	// +optional
	TagKey *string `json:"tagKey,omitempty" tf:"tag_key"`
	// Map of [string]string. At least one of the keys must match the value of default_value.
	ValuesToReadableStrings *map[string]string `json:"valuesToReadableStrings" tf:"values_to_readable_strings"`
}

type DashboardSpecSectionRowChartChartSetting struct {
	// deprecated
	// +optional
	AutoColumnTags *bool `json:"autoColumnTags,omitempty" tf:"auto_column_tags"`
	// deprecated
	// +optional
	ColumnTags *string `json:"columnTags,omitempty" tf:"column_tags"`
	// For the tabular view, a list of point tags to display when using the custom tag display mode
	// +optional
	CustomTags []string `json:"customTags,omitempty" tf:"custom_tags"`
	// Threshold (in seconds) for time delta between consecutive points in a series above which a dotted line will replace a solid line in line plots. Default: 60s
	// +optional
	ExpectedDataSpacing *int64 `json:"expectedDataSpacing,omitempty" tf:"expected_data_spacing"`
	// For a chart with a fixed legend, a list of statistics to display in the legend
	// +optional
	FixedLegendDisplayStats []string `json:"fixedLegendDisplayStats,omitempty" tf:"fixed_legend_display_stats"`
	// Whether to enable a fixed tabular legend adjacent to the chart
	// +optional
	FixedLegendEnabled *bool `json:"fixedLegendEnabled,omitempty" tf:"fixed_legend_enabled"`
	// Statistic to use for determining whether a series is displayed on the fixed legend = ['CURRENT', 'MEAN', 'MEDIAN', 'SUM', 'MIN', 'MAX', 'COUNT']
	// +optional
	FixedLegendFilterField *string `json:"fixedLegendFilterField,omitempty" tf:"fixed_legend_filter_field"`
	// Number of series to include in the fixed legend
	// +optional
	FixedLegendFilterLimit *int64 `json:"fixedLegendFilterLimit,omitempty" tf:"fixed_legend_filter_limit"`
	// Whether to display Top- or Bottom-ranked series in the fixed legend = ['TOP', 'BOTTOM']
	// +optional
	FixedLegendFilterSort *string `json:"fixedLegendFilterSort,omitempty" tf:"fixed_legend_filter_sort"`
	// deprecated
	// +optional
	FixedLegendHideLabel *bool `json:"fixedLegendHideLabel,omitempty" tf:"fixed_legend_hide_label"`
	// Where the fixed legend should be displayed with respect to the chart = ['RIGHT', 'TOP', 'LEFT', 'BOTTOM']
	// +optional
	FixedLegendPosition *string `json:"fixedLegendPosition,omitempty" tf:"fixed_legend_position"`
	// If true, the legend uses non-summarized stats instead of summarized
	// +optional
	FixedLegendUseRawStats *bool `json:"fixedLegendUseRawStats,omitempty" tf:"fixed_legend_use_raw_stats"`
	// For the tabular view, whether to group multi metrics into a single row by a common source. If false, each metric for each source is displayed in its own row. If true, multiple metrics for the same host will be displayed as different columns in the same row
	// +optional
	GroupBySource *bool `json:"groupBySource,omitempty" tf:"group_by_source"`
	// Whether to disable the display of the floating legend (but reenable it when the ctrl-key is pressed)
	// +optional
	InvertDynamicLegendHoverControl *bool `json:"invertDynamicLegendHoverControl,omitempty" tf:"invert_dynamic_legend_hover_control"`
	// Plot interpolation type. linear is default = ['linear', 'step-before', 'step-after', 'basis', 'cardinal', 'monotone']
	// +optional
	LineType *string `json:"lineType,omitempty" tf:"line_type"`
	// Max value of Y-axis. Set to null or leave blank for auto
	// +optional
	Max *float64 `json:"max,omitempty" tf:"max"`
	// Min value of Y-axis. Set to null or leave blank for auto
	// +optional
	Min *float64 `json:"min,omitempty" tf:"min"`
	// For the tabular view, how many point tags to display
	// +optional
	NumTags *int64 `json:"numTags,omitempty" tf:"num_tags"`
	// The Markdown content for a Markdown display, in plain text. Use this field instead of markdownContent
	// +optional
	PlainMarkdownContent *string `json:"plainMarkdownContent,omitempty" tf:"plain_markdown_content"`
	// For the tabular view, whether to display sources. Default: true
	// +optional
	ShowHosts *bool `json:"showHosts,omitempty" tf:"show_hosts"`
	// For the tabular view, whether to display labels. Default: true
	// +optional
	ShowLabels *bool `json:"showLabels,omitempty" tf:"show_labels"`
	// For the tabular view, whether to display raw values. Default: false
	// +optional
	ShowRawValues *bool `json:"showRawValues,omitempty" tf:"show_raw_values"`
	// For the tabular view, whether to display display values in descending order. Default: false
	// +optional
	SortValuesDescending *bool `json:"sortValuesDescending,omitempty" tf:"sort_values_descending"`
	// For the single stat view, the decimal precision of the displayed number
	// +optional
	SparklineDecimalPrecision *int64 `json:"sparklineDecimalPrecision,omitempty" tf:"sparkline_decimal_precision"`
	// For the single stat view, the color of the displayed text (when not dynamically determined). Values should be in rgba(, , ,  format
	// +optional
	SparklineDisplayColor *string `json:"sparklineDisplayColor,omitempty" tf:"sparkline_display_color"`
	// For the single stat view, the font size of the displayed text, in percent
	// +optional
	SparklineDisplayFontSize *string `json:"sparklineDisplayFontSize,omitempty" tf:"sparkline_display_font_size"`
	// For the single stat view, the horizontal position of the displayed text = ['MIDDLE', 'LEFT', 'RIGHT']
	// +optional
	SparklineDisplayHorizontalPosition *string `json:"sparklineDisplayHorizontalPosition,omitempty" tf:"sparkline_display_horizontal_position"`
	// For the single stat view, a string to append to the displayed text
	// +optional
	SparklineDisplayPostfix *string `json:"sparklineDisplayPostfix,omitempty" tf:"sparkline_display_postfix"`
	// For the single stat view, a string to add before the displayed text
	// +optional
	SparklineDisplayPrefix *string `json:"sparklineDisplayPrefix,omitempty" tf:"sparkline_display_prefix"`
	// For the single stat view, whether to display the name of the query or the value of query = ['VALUE', 'LABEL']
	// +optional
	SparklineDisplayValueType *string `json:"sparklineDisplayValueType,omitempty" tf:"sparkline_display_value_type"`
	// deprecated
	// +optional
	SparklineDisplayVerticalPosition *string `json:"sparklineDisplayVerticalPosition,omitempty" tf:"sparkline_display_vertical_position"`
	// For the single stat view, the color of the background fill. Values should be in rgba(, , ,  format
	// +optional
	SparklineFillColor *string `json:"sparklineFillColor,omitempty" tf:"sparkline_fill_color"`
	// For the single stat view, the color of the line. Values should be in rgba(, , ,  format
	// +optional
	SparklineLineColor *string `json:"sparklineLineColor,omitempty" tf:"sparkline_line_color"`
	// For the single stat view, a misleadingly named property. This determines whether the sparkline of the statistic is displayed in the chart BACKGROUND, BOTTOM, or NONE = ['BACKGROUND', 'BOTTOM', 'NONE']
	// +optional
	SparklineSize *string `json:"sparklineSize,omitempty" tf:"sparkline_size"`
	// For the single stat view, whether to apply dynamic color settings to the displayed TEXT or BACKGROUND = ['TEXT', 'BACKGROUND']
	// +optional
	SparklineValueColorMapApplyTo *string `json:"sparklineValueColorMapApplyTo,omitempty" tf:"sparkline_value_color_map_apply_to"`
	// For the single stat view, a list of colors that differing query values map to. Must contain one more element than sparklineValueColorMapValuesV2. Values should be in rgba(, , ,  format
	// +optional
	SparklineValueColorMapColors []string `json:"sparklineValueColorMapColors,omitempty" tf:"sparkline_value_color_map_colors"`
	// deprecated
	// +optional
	SparklineValueColorMapValues []int64 `json:"sparklineValueColorMapValues,omitempty" tf:"sparkline_value_color_map_values"`
	// deprecated
	// +optional
	SparklineValueColorMapValuesV2 []float64 `json:"sparklineValueColorMapValuesV2,omitempty" tf:"sparkline_value_color_map_values_v2"`
	// For the single stat view, a list of display text values that different query values map to. Must contain one more element than sparklineValueTextMapThresholds
	// +optional
	SparklineValueTextMapText []string `json:"sparklineValueTextMapText,omitempty" tf:"sparkline_value_text_map_text"`
	// For the single stat view, a list of threshold boundaries for mapping different query values to display text. Must contain one less element than sparklineValueTextMapText
	// +optional
	SparklineValueTextMapThresholds []float64 `json:"sparklineValueTextMapThresholds,omitempty" tf:"sparkline_value_text_map_thresholds"`
	// Type of stacked chart (applicable only if chart type is stacked). zero (default) means stacked from y=0. expand means Normalized from 0 to 1. wiggle means Minimize weighted changes. silhouette means to Center the Stream = ['zero', 'expand', 'wiggle', 'silhouette']
	// +optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type"`
	// For the tabular view, which mode to use to determine which point tags to display = ['all', 'top', 'custom']
	// +optional
	TagMode *string `json:"tagMode,omitempty" tf:"tag_mode"`
	// Fox x-y scatterplots, whether to color more recent points as darker than older points. Default: false
	// +optional
	TimeBasedColoring *bool `json:"timeBasedColoring,omitempty" tf:"time_based_coloring"`
	// Chart Type. 'line' refers to the Line Plot, 'scatter' to the Point Plot, 'stacked-area' to the Stacked Area plot, 'table' to the Tabular View, 'scatterploy-xy' to Scatter Plot, 'markdown-widget' to the Markdown display, and 'sparkline' to the Single Stat view = ['line', 'scatterplot', 'stacked-area', 'table', 'scatterplot-xy', 'markdown-widget', 'sparkline']
	Type *string `json:"type" tf:"type"`
	// Width, in minutes, of the time window to use for last windowing
	// +optional
	WindowSize *int64 `json:"windowSize,omitempty" tf:"window_size"`
	// For the tabular view, whether to use the full time window for the query or the last X minutes = ['full', 'last']
	// +optional
	Windowing *string `json:"windowing,omitempty" tf:"windowing"`
	// For x-y scatterplots, max value for X-axis. Set null for auto
	// +optional
	Xmax *float64 `json:"xmax,omitempty" tf:"xmax"`
	// For x-y scatterplots, min value for X-axis. Set null for auto
	// +optional
	Xmin *float64 `json:"xmin,omitempty" tf:"xmin"`
	// Default: false. Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI) ,
	// +optional
	Y0ScaleSiBy1024 *bool `json:"y0ScaleSiBy1024,omitempty" tf:"y0_scale_si_by_1024"`
	// Default: false. Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units
	// +optional
	Y0UnitAutoscaling *bool `json:"y0UnitAutoscaling,omitempty" tf:"y0_unit_autoscaling"`
	// Default: false. Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI)
	// +optional
	Y1ScaleSiBy1024 *bool `json:"y1ScaleSiBy1024,omitempty" tf:"y1_scale_si_by_1024"`
	// Default: false. Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units
	// +optional
	Y1UnitAutoscaling *bool `json:"y1UnitAutoscaling,omitempty" tf:"y1_unit_autoscaling"`
	// For plots with multiple Y-axes, units for right-side Y-axis
	// +optional
	Y1Units *string `json:"y1Units,omitempty" tf:"y1_units"`
	// For plots with multiple Y-axes, max value for right-side Y-axis. Set null for auto
	// +optional
	Y1max *float64 `json:"y1max,omitempty" tf:"y1max"`
	// For plots with multiple Y-axes, min value for right-side Y-axis. Set null for auto
	// +optional
	Y1min *float64 `json:"y1min,omitempty" tf:"y1min"`
	// For x-y scatterplots, max value for Y-axis. Set null for auto
	// +optional
	Ymax *float64 `json:"ymax,omitempty" tf:"ymax"`
	// For x-y scatterplots, min value for Y-axis. Set null for auto
	// +optional
	Ymin *float64 `json:"ymin,omitempty" tf:"ymin"`
}

type DashboardSpecSectionRowChartSource struct {
	// Whether to disabled the source from being displayed
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// Name of the Source
	Name *string `json:"name" tf:"name"`
	// Query for the Source
	Query *string `json:"query" tf:"query"`
	// Whether the query builder should be enabled
	// +optional
	QueryBuilderEnabled *bool `json:"queryBuilderEnabled,omitempty" tf:"query_builder_enabled"`
	// +optional
	ScatterPlotSource *string `json:"scatterPlotSource,omitempty" tf:"scatter_plot_source"`
	// Description of the source
	// +optional
	SourceDescription *string `json:"sourceDescription,omitempty" tf:"source_description"`
}

type DashboardSpecSectionRowChart struct {
	// Base of logarithmic scale, default is 0 for linear scale
	// +optional
	Base *int64 `json:"base,omitempty" tf:"base"`
	// +optional
	ChartAttribute *string `json:"chartAttribute,omitempty" tf:"chart_attribute"`
	// Chart settings. Defaults to line charts
	ChartSetting *DashboardSpecSectionRowChartChartSetting `json:"chartSetting" tf:"chart_setting"`
	// Description of the chart
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Name of the Chart
	Name *string `json:"name" tf:"name"`
	// A collection of Sources for a Chart
	Source []DashboardSpecSectionRowChartSource `json:"source" tf:"source"`
	// Summarization strategy for the chart. MEAN is default = ['MEAN', 'MEDIAN', 'MIN', 'MAX', 'SUM', 'COUNT', 'LAST', 'FIRST']
	Summarization *string `json:"summarization" tf:"summarization"`
	// Units of measurements for the chart
	Units *string `json:"units" tf:"units"`
}

type DashboardSpecSectionRow struct {
	// A collection of chart
	Chart []DashboardSpecSectionRowChart `json:"chart" tf:"chart"`
}

type DashboardSpecSection struct {
	// Name of the Sections
	Name *string `json:"name" tf:"name"`
	// Rows containing chart. Rows belong in Sections
	Row []DashboardSpecSectionRow `json:"row" tf:"row"`
}

type DashboardSpec struct {
	State *DashboardSpecResource `json:"state,omitempty" tf:"-"`

	Resource DashboardSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DashboardSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CanModify []string `json:"canModify,omitempty" tf:"can_modify"`
	// +optional
	CanView     []string `json:"canView,omitempty" tf:"can_view"`
	Description *string  `json:"description" tf:"description"`
	// +optional
	DisplayQueryParameters *bool `json:"displayQueryParameters,omitempty" tf:"display_query_parameters"`
	// +optional
	DisplaySectionTableOfContents *bool `json:"displaySectionTableOfContents,omitempty" tf:"display_section_table_of_contents"`
	// +optional
	EventFilterType *string `json:"eventFilterType,omitempty" tf:"event_filter_type"`
	Name            *string `json:"name" tf:"name"`
	// +optional
	ParameterDetails []DashboardSpecParameterDetails `json:"parameterDetails,omitempty" tf:"parameter_details"`
	// Sections of a Dashboard
	Section []DashboardSpecSection `json:"section" tf:"section"`
	Tags    []string               `json:"tags" tf:"tags"`
	Url     *string                `json:"url" tf:"url"`
}

type DashboardStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DashboardList is a list of Dashboards
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Dashboard CRD objects
	Items []Dashboard `json:"items,omitempty"`
}
