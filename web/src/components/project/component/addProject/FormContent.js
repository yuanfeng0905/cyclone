import React from 'react';
import { Field } from 'formik';
import PropTypes from 'prop-types';
import { Form, Input, Button } from 'antd';
import Quota from '@/components/public/quota';
import IntegrationSelect from './IntegrationSelect';
import MakeField from '@/components/public/makeField';

const { TextArea } = Input;
const InputField = MakeField(Input);
const TextareaField = MakeField(TextArea);

const FormContent = ({ history, handleSubmit, setFieldValue }) => {
  return (
    <Form layout={'horizontal'} onSubmit={handleSubmit}>
      <Field
        label={intl.get('name')}
        name="metadata.name"
        component={InputField}
        hasFeedback
        required
      />
      <Field
        label={intl.get('description')}
        name="metadata.description"
        component={TextareaField}
      />
      <Field
        label={intl.get('project.externalSystem')}
        name="spec.services"
        required
        component={IntegrationSelect}
      />
      <Field
        label={intl.get('allocation.quotaConfig')}
        name="spec.quota"
        component={Quota}
        onChange={value => {
          setFieldValue('spec.quota', value);
        }}
      />
      <div className="form-bottom-btn">
        <Button type="primary" htmlType="submit">
          {intl.get('operation.confirm')}
        </Button>
        <Button
          onClick={() => {
            history.push(`/project`);
          }}
        >
          {intl.get('operation.cancel')}
        </Button>
      </div>
    </Form>
  );
};

FormContent.propTypes = {
  history: PropTypes.object,
  handleSubmit: PropTypes.func,
  setFieldValue: PropTypes.func,
};

export default FormContent;
